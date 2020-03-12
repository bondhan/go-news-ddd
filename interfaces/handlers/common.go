package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
)

// InitCommon ..
func InitCommon(r *chi.Mux) {

	r.Get("/", hello)

	// healtcheck
	r.Get("/ping", ping)

}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome to go-ddd-news"))
}
