package handlers

import (
	"fmt"
	"net/http"

	"github.com/tomanagle/url-shortener/store"
)

type GetShortURLHandler struct {
	shortURLStore store.ShortURLStore
}

type GetShortURLHandlerParams struct {
	ShortURLStore store.ShortURLStore
}

func NewGetShortURLHandler(params GetShortURLHandlerParams) *GetShortURLHandler {
	return &GetShortURLHandler{
		shortURLStore: params.ShortURLStore,
	}
}

func (h *GetShortURLHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	slug := r.URL.Path[1:]

	fmt.Println("slug", slug)

	shortURL, err := h.shortURLStore.GetShortURLBySlug(slug)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if shortURL == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	http.Redirect(w, r, shortURL.Destination, http.StatusMovedPermanently)
}
