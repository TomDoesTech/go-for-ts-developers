package dbstore

import (
	"log/slog"

	"github.com/tomanagle/url-shortener/store"
)

type ShortURLStore struct {
	shortURLs []store.ShortURL
	logger    *slog.Logger
}

type NewShortURLStoreParams struct {
	Logger *slog.Logger
}

func NewShortURLStore(params NewShortURLStoreParams) *ShortURLStore {
	shortURLs := []store.ShortURL{}

	return &ShortURLStore{
		shortURLs: shortURLs,
		logger:    params.Logger,
	}
}

func (s *ShortURLStore) CreateShortURL(params store.CreateShortURLParams) (store.ShortURL, error) {

	shortURL := store.ShortURL{
		Destination: params.Destination,
		Slug:        params.Slug,
		ID:          len(s.shortURLs),
	}

	s.shortURLs = append(s.shortURLs, shortURL)

	s.logger.Info("short URL created", slog.Any("values", shortURL))

	return shortURL, nil
}

func (s *ShortURLStore) GetShortURLBySlug(slug string) (*store.ShortURL, error) {

	for _, shortURL := range s.shortURLs {
		if shortURL.Slug == slug {
			result := shortURL
			return &result, nil
		}
	}

	return nil, store.ErrShortURLNotFound
}
