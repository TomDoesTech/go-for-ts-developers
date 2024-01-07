package mock

import (
	"github.com/stretchr/testify/mock"
	"github.com/tomanagle/url-shortener/store"
)

type MockShortURLStore struct {
	mock.Mock
}

func (m *MockShortURLStore) CreateShortURL(params store.CreateShortURLParams) (store.ShortURL, error) {
	args := m.Called(params)
	return args.Get(0).(store.ShortURL), args.Error(1)
}

func (m *MockShortURLStore) GetShortURLBySlug(slug string) (*store.ShortURL, error) {
	args := m.Called(slug)
	return args.Get(0).(*store.ShortURL), args.Error(1)
}
