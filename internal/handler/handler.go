package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/kstsm/wb-shortener/internal/service"
	"net/http"
)

type HandlerI interface {
	NewRouter() http.Handler
}

type Handler struct {
	service service.ServiceI
}

func NewHandler(service service.ServiceI) HandlerI {
	return &Handler{
		service: service,
	}
}

func (h Handler) NewRouter() http.Handler {
	r := chi.NewRouter()

	r.Post("/shorten", h.shortenURLHandler)
	r.Get("/analytics/{shortURL}", h.getAnalytics)
	r.Get("/s/{shortURL}", h.redirect)

	r.Handle("/", http.FileServer(http.Dir("web/")))

	return r
}
