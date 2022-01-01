package infrastructure

import (
	"github.com/bondhan/godddnews/infrastructure/client"
	"github.com/bondhan/godddnews/interfaces/handlers/rest"
	"github.com/go-chi/chi"
)

func InitApplicationAndRouters(manager client.Manager) *chi.Mux {
	r := chi.NewRouter()

	utilitiesHandlers := rest.NewUtilitiesHandlers()
	newsHandlers := rest.NewNewsHandlers(manager)
	topicHandlers := rest.NewTopicHandlers(manager)
	tagHandlers := rest.NewTagHandlers(manager)

	rest.InitUtilitiesRouter(r, utilitiesHandlers)
	rest.InitNewsRouter(r, newsHandlers)
	rest.InitTopicRouter(r, topicHandlers)
	rest.InitTagRouter(r, tagHandlers)

	return r
}
