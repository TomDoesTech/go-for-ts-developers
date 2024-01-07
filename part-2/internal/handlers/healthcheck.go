package handlers

import "net/http"

type HealthcheckHandler struct {
}

func NewHealthcheckHandler() *HealthcheckHandler {
	return &HealthcheckHandler{}
}

func (h *HealthcheckHandler) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
