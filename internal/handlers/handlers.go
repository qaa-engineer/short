package handlers

import (
	"io"
	"net/http"
	"net/url"

	"github.com/go-chi/chi/v5"
	"github.com/qaa-engineer/short/internal/hasher"
	"github.com/qaa-engineer/short/internal/storage"
)

type URLShortenerHandler struct {
	URLRepository storage.URLRepository
}

func NewURLShortenerHandler() *URLShortenerHandler {
	return &URLShortenerHandler{
		URLRepository: storage.NewURLStorage(),
	}
}

func (handler *URLShortenerHandler) PostHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	w.Header().Set("Content-Type", "text/plain")

	if err != nil || len(body) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := url.Parse(string(body))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	fullLink := res.String()
	shortLink, err := hasher.GetShortLink(fullLink)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	handler.URLRepository.AddURL(shortLink, fullLink)

	w.WriteHeader(http.StatusCreated)
	host := r.Host

	_, err = w.Write([]byte("http://" + host + "/" + shortLink))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (handler *URLShortenerHandler) GetHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	url, ok := handler.URLRepository.GetURL(id)

	w.Header().Set("Content-Type", "text/plain")

	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Location", url)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
