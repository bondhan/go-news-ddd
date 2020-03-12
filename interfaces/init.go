package interfaces

import (
	"net/http"

	"github.com/bondhan/godddnews/interfaces/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

func Init(port string) {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	handlers.InitCommon(r)
	handlers.InitNewsHandler(r)
	handlers.InitTopicHandler(r)
	handlers.InitTagHandler(r)

	http.ListenAndServe(":"+port, r)
}
