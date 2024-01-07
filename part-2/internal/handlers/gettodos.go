package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/tomdoestech/go-for-ts-devs/internal/store"
)

type GetTodosHandler struct {
	todos *[]store.Todo
}

type GetTodosHandlerParams struct {
	Todos *[]store.Todo
}

func NewGetTodosHandler(params GetTodosHandlerParams) *GetTodosHandler {
	return &GetTodosHandler{
		todos: params.Todos,
	}
}

func (h *GetTodosHandler) ServerHTTP(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(h.todos)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
