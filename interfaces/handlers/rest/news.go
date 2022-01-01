package rest

import (
	"github.com/bondhan/godddnews/domain/constants"
	"github.com/bondhan/godddnews/infrastructure/client"
	"github.com/go-chi/chi"
	"net/http"
)

// InitNewsRouter .
func InitNewsRouter(r *chi.Mux, n NewsHandlers) http.Handler {

	r.Route("/news", func(r chi.Router) {
		r.Get("/", n.getAllNews)
		r.Post("/", n.createNews)
	})

	return r
}

type news struct {
	manager client.Manager
}

type NewsHandlers interface {
	getAllNews(w http.ResponseWriter, r *http.Request)
	createNews(w http.ResponseWriter, r *http.Request)
}

func NewNewsHandlers(mgr client.Manager) NewsHandlers {
	return &news{manager: mgr}
}

func (n *news) getAllNews(w http.ResponseWriter, r *http.Request) {
	respClient := n.manager.GetObject(constants.RespondClient).(client.ResponseClient)

	respClient.JSON(w, r, http.StatusOK, "getAllNews")
}

func (n *news) createNews(w http.ResponseWriter, r *http.Request) {
	respClient := n.manager.GetObject(constants.RespondClient).(client.ResponseClient)
	respClient.JSON(w, r, http.StatusOK, "createNews")
}
