package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/qaa-engineer/short/internal/handlers"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.AllowContentType("text/plain"))

	urlShortenerHandler := handlers.NewURLShortenerHandler()

	r.Post("/", urlShortenerHandler.PostHandler)
	r.Get("/{id}", urlShortenerHandler.GetHandler)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
