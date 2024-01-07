package handlers

import "net/http"

type HelthcheckHandler struct{}

func NewHealthHandler() *HelthcheckHandler {
	return &HelthcheckHandler{}
}

func (h *HelthcheckHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
