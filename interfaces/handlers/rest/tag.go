package rest

import (
	"github.com/bondhan/godddnews/infrastructure/client"
	"github.com/go-chi/chi"
	"net/http"
)

// InitTagRouter .
func InitTagRouter(r *chi.Mux, t TagHandlers) http.Handler {
	r.Route("/api/v1/tag", func(r chi.Router) {
		r.Get("/", t.getAllTag)
		r.Post("/", t.createTag)

	})

	return r
}

type tag struct {
	manager client.Manager
}

type TagHandlers interface {
	getAllTag(w http.ResponseWriter, r *http.Request)
	createTag(w http.ResponseWriter, r *http.Request)
}

func NewTagHandlers(mgr client.Manager) TagHandlers {
	return &tag{manager: mgr}
}

func (t *tag) getAllTag(w http.ResponseWriter, r *http.Request) {
}

func (t *tag) createTag(w http.ResponseWriter, r *http.Request) {

}
