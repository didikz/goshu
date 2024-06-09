package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"path"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
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
	// slug := chi.URLParam(r, "slug")
	w.Header().Add("Cache-Control", "no-cache")
	http.Redirect(w, r, "https://didiktrisusanto.dev", http.StatusMovedPermanently)
	return nil
}
