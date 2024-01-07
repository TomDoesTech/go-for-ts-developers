package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tomanagle/url-shortener/store"
	storemock "github.com/tomanagle/url-shortener/store/mock"
)

func TestCreateShortURL(t *testing.T) {

	testCases := []struct {
		name                 string
		createMockResult     store.ShortURL
		expectedStatusCode   int
		expectedBody         []byte
		payload              string
		expectedCreateParams store.CreateShortURLParams
	}{
		{
			name:               "success",
			expectedStatusCode: http.StatusCreated,
			createMockResult: store.ShortURL{
				Slug:        "abcdef",
				Destination: "https://www.google.com",
			},
			expectedBody: []byte(`{"destination":"https://www.google.com", "id":0, "slug":"abcdef"}`),
			expectedCreateParams: store.CreateShortURLParams{
				Destination: "https://www.google.com",
				Slug:        "1234",
			},
			payload: `{"destination": "https://www.google.com"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			mockShortURLStore := &storemock.MockShortURLStore{}

			mockShortURLStore.On("CreateShortURL", tc.expectedCreateParams).Return(tc.createMockResult, nil)

			handler := NewCreateShortURLHandler(CreateShortURLHandlerParams{
				ShortURLStore: mockShortURLStore,
				GenerateSlug: func() string {
					return "1234"
				},
			})

			request := httptest.NewRequest("POST", "/shorturl", strings.NewReader(tc.payload))
			responseRecorder := httptest.NewRecorder()

			handler.ServeHTTP(responseRecorder, request)

			response := responseRecorder.Result()
			defer response.Body.Close()
			body, err := io.ReadAll(response.Body)
			assert.NoError(err)

			assert.Equal(tc.expectedStatusCode, response.StatusCode)
			assert.JSONEq(string(tc.expectedBody), string(body))

			mockShortURLStore.AssertExpectations(t)
		})
	}

}
