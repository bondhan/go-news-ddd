package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/bondhan/godddnews/application"
	"github.com/bondhan/godddnews/domain"
	"github.com/bondhan/godddnews/interfaces/respond"
	"github.com/bondhan/godddnews/internal/manager"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

// InitTagHandler ..
func InitTagHandler(r *chi.Mux) {
	r.Route("/api/v1/tag", func(r chi.Router) {
		r.Get("/", getAllTag)
		r.Post("/", createTag)

		r.Route("/{tagID}", func(r chi.Router) {
			r.Get("/", getTagByID)
			r.Put("/", updateTagByID)
			r.Delete("/", deleteTagByID)
		})

		r.Get("/{tagSlug:[a-z-]+}", getTagBySlug)
		r.Put("/{tagSlug:[a-z-]+}", updateTagBySlug)
		r.Delete("/{tagSlug:[a-z-]+}", deleteTagBySlug)
	})
}

// getAllNews ..
func getAllTag(w http.ResponseWriter, r *http.Request) {
	manager.GetContainer().Invoke(func(TagApp application.TagApp) {
		tags, err := TagApp.GetAllTags()
		if err != nil {
			respond.Error(w, http.StatusNotFound, fmt.Errorf("%s", err.Error()))
		}

		respond.JSON(w, http.StatusOK, tags)
	})
}

func createTag(w http.ResponseWriter, r *http.Request) {
	manager.GetContainer().Invoke(func(TagApp application.TagApp) {
		decoder := json.NewDecoder(r.Body)
		var p domain.Tag
		if err := decoder.Decode(&p); err != nil {
			respond.Error(w, http.StatusBadRequest, fmt.Errorf("wrong JSON parameter: %s", err.Error()))
			return
		}

		err := TagApp.AddTag(p)
		if err != nil {
			respond.Error(w, http.StatusConflict, fmt.Errorf("%s", err.Error()))
			return
		}

		respond.Success(w, http.StatusCreated, "Success")
	})
}

// getTagBySlug ..
func getTagBySlug(w http.ResponseWriter, r *http.Request) {
	manager.GetContainer().Invoke(func(TagApp application.TagApp) {

		logrus.Debug("getTagBySlug")
		param := chi.URLParam(r, "tagSlug")

		news, err := TagApp.GetTagBySlug(param)
		if err != nil {
			respond.Error(w, http.StatusNotFound, fmt.Errorf("%s", err.Error()))
			return
		}

		respond.JSON(w, http.StatusOK, news)
	})
}

func updateTagBySlug(w http.ResponseWriter, r *http.Request) {
	manager.GetContainer().Invoke(func(TagApp application.TagApp) {
		decoder := json.NewDecoder(r.Body)
		var p domain.Tag
		if err := decoder.Decode(&p); err != nil {
			respond.Error(w, http.StatusBadRequest, fmt.Errorf("wrong JSON parameter: %s", err.Error()))
			return
		}

		logrus.Debug("updateTagBySlug")
		param := chi.URLParam(r, "tagSlug")

		err := TagApp.UpdateTagBySlug(p, param)
		if err != nil {
			respond.Error(w, http.StatusConflict, fmt.Errorf("%s", err.Error()))
			return
		}

		respond.Success(w, http.StatusCreated, "Success")
	})
}

// deleteTagBySlug ..
func deleteTagBySlug(w http.ResponseWriter, r *http.Request) {
	manager.GetContainer().Invoke(func(TagApp application.TagApp) {

		logrus.Debug("deleteTagBySlug")
		param := chi.URLParam(r, "tagSlug")

		err := TagApp.DeleteTagBySlug(param)
		if err != nil {
			respond.Error(w, http.StatusConflict, fmt.Errorf("%s", err.Error()))
			return
		}

		respond.Success(w, http.StatusOK, "Success")
	})
}

// getTagByID ..
func getTagByID(w http.ResponseWriter, r *http.Request) {
	manager.GetContainer().Invoke(func(TagApp application.TagApp) {

		logrus.Debug("getTagByID")
		param := chi.URLParam(r, "tagID")

		p, err := strconv.Atoi(param)
		if err != nil {
			respond.Error(w, http.StatusNotFound, fmt.Errorf("Data Not Found"))
			return
		}

		news, err := TagApp.GetTagByID(uint(p))
		if err != nil {
			respond.Error(w, http.StatusNotFound, fmt.Errorf("%s", err.Error()))
			return
		}

		respond.JSON(w, http.StatusOK, news)
	})
}

func updateTagByID(w http.ResponseWriter, r *http.Request) {
	manager.GetContainer().Invoke(func(TagApp application.TagApp) {
		decoder := json.NewDecoder(r.Body)
		var p domain.Tag
		if err := decoder.Decode(&p); err != nil {
			respond.Error(w, http.StatusBadRequest, fmt.Errorf("wrong JSON parameter: %s", err.Error()))
			return
		}

		logrus.Debug("updateTagByID")
		param := chi.URLParam(r, "tagID")

		id, err := strconv.Atoi(param)
		if err != nil {
			respond.Error(w, http.StatusNotFound, fmt.Errorf("Data Not Found"))
			return
		}

		err = TagApp.UpdateTagByID(p, uint(id))
		if err != nil {
			respond.Error(w, http.StatusConflict, fmt.Errorf("%s", err.Error()))
			return
		}

		respond.Success(w, http.StatusCreated, "Success")
	})
}

// deleteTagByID ..
func deleteTagByID(w http.ResponseWriter, r *http.Request) {
	manager.GetContainer().Invoke(func(TagApp application.TagApp) {

		logrus.Debug("deleteTagByID")
		param := chi.URLParam(r, "tagID")

		id, err := strconv.Atoi(param)
		if err != nil {
			respond.Error(w, http.StatusNotFound, fmt.Errorf("Data Not Found"))
			return
		}

		err = TagApp.DeleteTagByID(uint(id))
		if err != nil {
			respond.Error(w, http.StatusConflict, fmt.Errorf("%s", err.Error()))
			return
		}

		respond.Success(w, http.StatusOK, "Success")
	})
}
