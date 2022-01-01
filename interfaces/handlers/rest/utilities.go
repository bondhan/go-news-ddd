package rest

import (
	"net/http"

	"github.com/go-chi/chi"
)

func InitUtilitiesRouter(r *chi.Mux, u UtilitiesHandlers) {
	r.Get("/", u.hello)
	r.Get("/ping", u.ping)
}

type utilities struct {
}

type UtilitiesHandlers interface {
	ping(w http.ResponseWriter, r *http.Request)
	hello(w http.ResponseWriter, r *http.Request)
}

func NewUtilitiesHandlers() UtilitiesHandlers {
	return &utilities{}
}

func (u utilities) ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func (u utilities) hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome to go-ddd-news"))
}
