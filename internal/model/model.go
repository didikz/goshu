package model

type GeneralApiError struct {
	Error string `json:"error"`
}

type IndexResponse struct {
	Message string `json:"message"`
}

type ShortenUrlResult struct {
	Url string `json:"url"`
}

type Url struct {
	ID          uint64  `db:"id"`
	Slug        string  `db:"slug"`
	OriginalUrl string  `db:"original_url"`
	UserID      *string `db:"user_id"`
	VisitCount  int     `db:"visit_count"`
	CreatedAt   string  `db:"created_at"`
	UpdateAt    *string `db:"update_at"`
	DeletedAt   *string `db:"deleted_at"`
}
