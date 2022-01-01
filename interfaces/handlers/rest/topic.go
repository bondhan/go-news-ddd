package rest

import (
	"github.com/bondhan/godddnews/infrastructure/client"
	"github.com/go-chi/chi"
	"net/http"
)

// InitTopicRouter .
func InitTopicRouter(r *chi.Mux, t TopicHandler) http.Handler {
	r.Route("/api/v1/topic", func(r chi.Router) {
		r.Get("/", t.getAllTopic)
		r.Post("/", t.createTopic)
	})

	return r
}

type topic struct {
	manager client.Manager
}

type TopicHandler interface {
	getAllTopic(w http.ResponseWriter, r *http.Request)
	createTopic(w http.ResponseWriter, r *http.Request)
}

func NewTopicHandlers(mgr client.Manager) TopicHandler {
	return &topic{manager: mgr}
}

func (t *topic) getAllTopic(w http.ResponseWriter, r *http.Request) {
}

func (t *topic) createTopic(w http.ResponseWriter, r *http.Request) {
}
