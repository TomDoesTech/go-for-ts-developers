package store

import "errors"

var ErrShortURLNotFound = errors.New("short URL not found")

type ShortURL struct {
	ID          int    `json:"id" bson:"shortURLID"`
	Destination string `json:"destination" bson:"dest"`
	Slug        string `json:"slug" bson:"slug"`
}

type CreateShortURLParams struct {
	Destination string
	Slug        string
}

type ShortURLStore interface {
	CreateShortURL(params CreateShortURLParams) (ShortURL, error)
	GetShortURLBySlug(slug string) (*ShortURL, error)
}
