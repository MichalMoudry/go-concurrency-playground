package transport

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Handler struct {
	Port int
	Mux  *chi.Mux
}

func NewHandler(port int) *Handler {
	handler := &Handler{
		Port: port,
		Mux:  chi.NewRouter(),
	}
	handler.Mux.Use(middleware.Logger)

	handler.Mux.Get("/health", health)
	return handler
}

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}
