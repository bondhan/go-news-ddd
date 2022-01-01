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

	routers := chi.NewRouter()
	rest.InitNewsRouter(routers, newsHandlers)
	rest.InitTopicRouter(routers, topicHandlers)
	rest.InitTagRouter(routers, tagHandlers)

	r.Mount("/api/v1", routers)
	r.NotFoundHandler()

	return r
}
