package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type handleFunc func(http.ResponseWriter, *http.Request) error

func WriteResponseJSON(w http.ResponseWriter, httpStatus int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	return json.NewEncoder(w).Encode(v)
}
func makeHttpHandleFunc(h handleFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			WriteResponseJSON(w, http.StatusBadRequest, GeneralApiError{Error: err.Error()})
		}
	}
}

type Server struct {
	listenAddr string
	db         *sqlx.DB
	redis      *redis.Client
}

func NewServer(listenAddr string, db *sqlx.DB, redis *redis.Client) *Server {
	return &Server{
		listenAddr: listenAddr,
		db:         db,
		redis:      redis,
	}
}

func (s *Server) Run() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.HandleFunc("/", makeHttpHandleFunc(s.handleIndex))
	router.HandleFunc("/{slug}", makeHttpHandleFunc(s.handleSlugRedirect))

	log.Println("Server running at", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}

func (s *Server) handleIndex(w http.ResponseWriter, r *http.Request) error {
	filepath := path.Join("views", "index.html")
	templ, err := template.ParseFiles(filepath)
	if err != nil {
		return err
	}
	return templ.Execute(w, nil)
}

func (s *Server) handleSlugRedirect(w http.ResponseWriter, r *http.Request) error {
	slug := chi.URLParam(r, "slug")
	w.Header().Add("Cache-Control", "no-cache")
	if slug != "" {
		url := Url{}
		cached := s.redis.Get(r.Context(), fmt.Sprintf("slug:%s", slug))
		if cached.Err() == nil && cached.Val() != "" {
			bytes, _ := cached.Bytes()
			_ = json.Unmarshal(bytes, &url)
			http.Redirect(w, r, url.OriginalUrl, http.StatusMovedPermanently)
			return nil
		}

		err := s.db.Get(&url, "SELECT id, slug, original_url, deleted_at FROM urls WHERE urls.slug = $1 LIMIT 1", slug)
		if err != nil {
			if err == sql.ErrNoRows {
				return fmt.Errorf("%s", "URL not found")
			}
			return err
		}

		if url.DeletedAt != nil {
			return fmt.Errorf("%s", "URL not found")
		}

		m, _ := json.Marshal(url)
		_ = s.redis.Set(r.Context(), fmt.Sprintf("slug:%s", slug), string(m), (time.Minute * 5))
		http.Redirect(w, r, url.OriginalUrl, http.StatusMovedPermanently)
	}
	return nil
}
