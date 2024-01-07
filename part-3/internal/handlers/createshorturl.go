package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/tomanagle/url-shortener/store"
)

type CreateShortURLHandler struct {
	shortURLStore store.ShortURLStore
	generateSlug  func() string
}

type CreateShortURLHandlerParams struct {
	ShortURLStore store.ShortURLStore
	GenerateSlug  func() string
}

func NewCreateShortURLHandler(params CreateShortURLHandlerParams) *CreateShortURLHandler {
	return &CreateShortURLHandler{
		shortURLStore: params.ShortURLStore,
		generateSlug:  params.GenerateSlug,
	}
}

func (h *CreateShortURLHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var requestData = struct {
		Destination string `json:"destination"`
	}{}

	err := json.NewDecoder(r.Body).Decode(&requestData)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	slug := h.generateSlug()

	createdShortURL, err := h.shortURLStore.CreateShortURL(store.CreateShortURLParams{
		Destination: requestData.Destination,
		Slug:        slug,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(createdShortURL)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
