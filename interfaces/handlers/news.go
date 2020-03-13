package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/bondhan/godddnews/application"
	"github.com/bondhan/godddnews/application/view"
	"github.com/bondhan/godddnews/infrastructure/manager"
	"github.com/bondhan/godddnews/interfaces/respond"
	"github.com/bondhan/godddnews/internal/utils"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

// InitNewsHandler ..
func InitNewsHandler(r *chi.Mux) {

	r.Route("/api/v1/news", func(r chi.Router) {
		r.Get("/", getAllNews)
		r.Post("/", createNews)

		r.Route("/{newsID}", func(r chi.Router) {
			r.Get("/", getNewsByID)
			r.Put("/", updateNewsByID)
			r.Delete("/", deleteNewsByID)
		})

		r.Get("/{newsSlug:[a-z-]+}", getNewsBySlug)
		r.Put("/{newsSlug:[a-z-]+}", updateNewsBySlug)
		r.Delete("/{newsSlug:[a-z-]+}", deleteNewsBySlug)

		r.Get("/topic/{topicSlug:[a-z-]+}", getNewsByTopicSlug)
		r.Get("/tag/{tagSlug:[a-z-]+}", getNewsByTagSlug)
	})
}

func getAllNews(w http.ResponseWriter, r *http.Request) {
	manager.GetContainer().Invoke(func(newsApp application.NewsApp) {
		queryValues := r.URL.Query()

		var newsView view.NewsView
		var err error
		newsView, err = newsApp.GetAllNews(queryValues)
		if err != nil {
			respond.Error(w, http.StatusNotFound, fmt.Errorf("%s", err.Error()))
		}

		respond.JSON(w, http.StatusOK, newsView)
	})
}

func createNews(w http.ResponseWriter, r *http.Request) {
	manager.GetContainer().Invoke(func(newsApp application.NewsApp) {
		decoder := json.NewDecoder(r.Body)
		var p view.NewsData
		if err := decoder.Decode(&p); err != nil {
			respond.Error(w, http.StatusBadRequest, fmt.Errorf("wrong JSON parameter: %s", err.Error()))
			return
		}

		err := utils.ValidateModels(p)
		if err != nil {
			respond.Error(w, http.StatusBadRequest, fmt.Errorf("%s", err.Error()))
			return
		}

		err = utils.ValidateSlug(p.Slug)
		if err != nil {
			respond.Error(w, http.StatusBadRequest, fmt.Errorf("%s", err.Error()))
			return
		}

		err = newsApp.AddNews(p)
		if err != nil {
			respond.Error(w, http.StatusConflict, fmt.Errorf("%s", err.Error()))
			return
		}

		respond.Success(w, http.StatusCreated, "Success")
	})
}

func getNewsBySlug(w http.ResponseWriter, r *http.Request) {
	manager.GetContainer().Invoke(func(newsApp application.NewsApp) {

		param := chi.URLParam(r, "newsSlug")

		news, err := newsApp.GetNewsBySlug(param)
		if err != nil {
			respond.Error(w, http.StatusNotFound, fmt.Errorf("%s", err.Error()))
			return
		}

		respond.JSON(w, http.StatusOK, news)
	})
}

func updateNewsBySlug(w http.ResponseWriter, r *http.Request) {
	manager.GetContainer().Invoke(func(newsApp application.NewsApp) {
		decoder := json.NewDecoder(r.Body)
		var p view.NewsData
		if err := decoder.Decode(&p); err != nil {
			respond.Error(w, http.StatusBadRequest, fmt.Errorf("wrong JSON parameter: %s", err.Error()))
			return
		}

		err := utils.ValidateModels(p)
		if err != nil {
			respond.Error(w, http.StatusBadRequest, fmt.Errorf("%s", err.Error()))
			return
		}

		err = utils.ValidateSlug(p.Slug)
		if err != nil {
			respond.Error(w, http.StatusBadRequest, fmt.Errorf("%s", err.Error()))
			return
		}

		param := chi.URLParam(r, "newsSlug")
		err = newsApp.UpdateNewsBySlug(p, param)
		if err != nil {
			respond.Error(w, http.StatusConflict, fmt.Errorf("%s", err.Error()))
			return
		}

		respond.Success(w, http.StatusCreated, "Success")
	})
}

func deleteNewsBySlug(w http.ResponseWriter, r *http.Request) {
	manager.GetContainer().Invoke(func(newsApp application.NewsApp) {

		param := chi.URLParam(r, "newsSlug")

		err := newsApp.DeleteNewsBySlug(param)
		if err != nil {
			respond.Error(w, http.StatusConflict, fmt.Errorf("%s", err.Error()))
			return
		}

		respond.Success(w, http.StatusOK, "Success")
	})
}

func getNewsByID(w http.ResponseWriter, r *http.Request) {
	manager.GetContainer().Invoke(func(newsApp application.NewsApp) {

		logrus.Debug("getNewsByID")
		param := chi.URLParam(r, "newsID")

		p, err := strconv.Atoi(param)
		if err != nil {
			respond.Error(w, http.StatusNotFound, fmt.Errorf("Data Not Found"))
			return
		}

		news, err := newsApp.GetNewsByID(uint(p))
		if err != nil {
			respond.Error(w, http.StatusNotFound, fmt.Errorf("%s", err.Error()))
			return
		}

		respond.JSON(w, http.StatusOK, news)
	})
}

func updateNewsByID(w http.ResponseWriter, r *http.Request) {
	manager.GetContainer().Invoke(func(newsApp application.NewsApp) {
		decoder := json.NewDecoder(r.Body)
		var p view.NewsData
		if err := decoder.Decode(&p); err != nil {
			respond.Error(w, http.StatusBadRequest, fmt.Errorf("wrong JSON parameter: %s", err.Error()))
			return
		}

		logrus.Debug("updateNewsByID")

		err := utils.ValidateModels(p)
		if err != nil {
			respond.Error(w, http.StatusBadRequest, fmt.Errorf("%s", err.Error()))
			return
		}

		err = utils.ValidateSlug(p.Slug)
		if err != nil {
			respond.Error(w, http.StatusBadRequest, fmt.Errorf("%s", err.Error()))
			return
		}

		param := chi.URLParam(r, "newsID")
		id, err := strconv.Atoi(param)
		if err != nil {
			respond.Error(w, http.StatusNotFound, fmt.Errorf("Data Not Found"))
			return
		}

		err = newsApp.UpdateNewsByID(p, uint(id))
		if err != nil {
			respond.Error(w, http.StatusConflict, fmt.Errorf("%s", err.Error()))
			return
		}

		respond.Success(w, http.StatusCreated, "Success")
	})
}

func deleteNewsByID(w http.ResponseWriter, r *http.Request) {
	manager.GetContainer().Invoke(func(newsApp application.NewsApp) {

		logrus.Debug("deleteNewsByID")
		param := chi.URLParam(r, "newsID")

		id, err := strconv.Atoi(param)
		if err != nil {
			respond.Error(w, http.StatusNotFound, fmt.Errorf("Data Not Found"))
			return
		}

		err = newsApp.DeleteNewsByID(uint(id))
		if err != nil {
			respond.Error(w, http.StatusConflict, fmt.Errorf("%s", err.Error()))
			return
		}

		respond.Success(w, http.StatusOK, "Success")
	})
}

func getNewsByTopicSlug(w http.ResponseWriter, r *http.Request) {
	manager.GetContainer().Invoke(func(newsApp application.NewsApp) {

		logrus.Debug("getNewsByTopicSlug")
		param := chi.URLParam(r, "topicSlug")

		news, err := newsApp.GetNewsByTopicSlug(param)
		if err != nil {
			respond.Error(w, http.StatusNotFound, fmt.Errorf("%s", err.Error()))
			return
		}

		respond.JSON(w, http.StatusOK, news)
	})
}

func getNewsByTagSlug(w http.ResponseWriter, r *http.Request) {
	manager.GetContainer().Invoke(func(newsApp application.NewsApp) {

		logrus.Debug("getNewsByTagSlug")
		param := chi.URLParam(r, "tagSlug")

		news, err := newsApp.GetNewsByTagSlug(param)
		if err != nil {
			respond.Error(w, http.StatusNotFound, fmt.Errorf("%s", err.Error()))
			return
		}

		respond.JSON(w, http.StatusOK, news)
	})
}
