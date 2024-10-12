package internal

type GeneralApiError struct {
	Error string `json:"error"`
}

type IndexResponse struct {
	Message string `json:"message"`
}
