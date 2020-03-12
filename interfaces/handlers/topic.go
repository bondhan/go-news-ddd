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

// InitTopicHandler ..
func InitTopicHandler(r *chi.Mux) {
	r.Route("/api/v1/topic", func(r chi.Router) {
		r.Get("/", getAllTopic)
		r.Post("/", createTopic)

		r.Route("/{topicID}", func(r chi.Router) {
			r.Get("/", getTopicByID)
			r.Put("/", updateTopicByID)
			r.Delete("/", deleteTopicByID)
		})

		r.Get("/{topicSlug:[a-z-]+}", getTopicBySlug)
		r.Put("/{topicSlug:[a-z-]+}", updateTopicBySlug)
		r.Delete("/{topicSlug:[a-z-]+}", deleteTopicBySlug)
	})
}

// getAllNews ..
func getAllTopic(w http.ResponseWriter, r *http.Request) {
	manager.GetContainer().Invoke(func(TopicApp application.TopicApp) {
		topics, err := TopicApp.GetAllTopics()
		if err != nil {
			respond.Error(w, http.StatusNotFound, fmt.Errorf("%s", err.Error()))
		}

		respond.JSON(w, http.StatusOK, topics)
	})
}

func createTopic(w http.ResponseWriter, r *http.Request) {
	manager.GetContainer().Invoke(func(TopicApp application.TopicApp) {
		decoder := json.NewDecoder(r.Body)
		var p domain.Topic
		if err := decoder.Decode(&p); err != nil {
			respond.Error(w, http.StatusBadRequest, fmt.Errorf("wrong JSON parameter: %s", err.Error()))
			return
		}

		err := TopicApp.AddTopic(p)
		if err != nil {
			respond.Error(w, http.StatusConflict, fmt.Errorf("%s", err.Error()))
			return
		}

		respond.Success(w, http.StatusCreated, "Success")
	})
}

// getTopicBySlug ..
func getTopicBySlug(w http.ResponseWriter, r *http.Request) {
	manager.GetContainer().Invoke(func(TopicApp application.TopicApp) {

		logrus.Debug("getTopicBySlug")
		param := chi.URLParam(r, "topicSlug")

		news, err := TopicApp.GetTopicBySlug(param)
		if err != nil {
			respond.Error(w, http.StatusNotFound, fmt.Errorf("%s", err.Error()))
			return
		}

		respond.JSON(w, http.StatusOK, news)
	})
}

func updateTopicBySlug(w http.ResponseWriter, r *http.Request) {
	manager.GetContainer().Invoke(func(TopicApp application.TopicApp) {
		decoder := json.NewDecoder(r.Body)
		var p domain.Topic
		if err := decoder.Decode(&p); err != nil {
			respond.Error(w, http.StatusBadRequest, fmt.Errorf("wrong JSON parameter: %s", err.Error()))
			return
		}

		logrus.Debug("updateTopicBySlug")
		param := chi.URLParam(r, "topicSlug")

		err := TopicApp.UpdateTopicBySlug(p, param)
		if err != nil {
			respond.Error(w, http.StatusConflict, fmt.Errorf("%s", err.Error()))
			return
		}

		respond.Success(w, http.StatusCreated, "Success")
	})
}

// deleteTopicBySlug ..
func deleteTopicBySlug(w http.ResponseWriter, r *http.Request) {
	manager.GetContainer().Invoke(func(TopicApp application.TopicApp) {

		logrus.Debug("deleteTopicBySlug")
		param := chi.URLParam(r, "topicSlug")

		err := TopicApp.DeleteTopicBySlug(param)
		if err != nil {
			respond.Error(w, http.StatusConflict, fmt.Errorf("%s", err.Error()))
			return
		}

		respond.Success(w, http.StatusOK, "Success")
	})
}

// getTopicByID ..
func getTopicByID(w http.ResponseWriter, r *http.Request) {
	manager.GetContainer().Invoke(func(TopicApp application.TopicApp) {

		logrus.Debug("getTopicByID")
		param := chi.URLParam(r, "topicID")

		p, err := strconv.Atoi(param)
		if err != nil {
			respond.Error(w, http.StatusNotFound, fmt.Errorf("Data Not Found"))
			return
		}

		news, err := TopicApp.GetTopicByID(uint(p))
		if err != nil {
			respond.Error(w, http.StatusNotFound, fmt.Errorf("%s", err.Error()))
			return
		}

		respond.JSON(w, http.StatusOK, news)
	})
}

func updateTopicByID(w http.ResponseWriter, r *http.Request) {
	manager.GetContainer().Invoke(func(TopicApp application.TopicApp) {
		decoder := json.NewDecoder(r.Body)
		var p domain.Topic
		if err := decoder.Decode(&p); err != nil {
			respond.Error(w, http.StatusBadRequest, fmt.Errorf("wrong JSON parameter: %s", err.Error()))
			return
		}

		logrus.Debug("updateTopicByID")
		param := chi.URLParam(r, "topicID")

		id, err := strconv.Atoi(param)
		if err != nil {
			respond.Error(w, http.StatusNotFound, fmt.Errorf("Data Not Found"))
			return
		}

		err = TopicApp.UpdateTopicByID(p, uint(id))
		if err != nil {
			respond.Error(w, http.StatusConflict, fmt.Errorf("%s", err.Error()))
			return
		}

		respond.Success(w, http.StatusCreated, "Success")
	})
}

// deleteTopicByID ..
func deleteTopicByID(w http.ResponseWriter, r *http.Request) {
	manager.GetContainer().Invoke(func(TopicApp application.TopicApp) {

		logrus.Debug("deleteTopicByID")
		param := chi.URLParam(r, "topicID")

		id, err := strconv.Atoi(param)
		if err != nil {
			respond.Error(w, http.StatusNotFound, fmt.Errorf("Data Not Found"))
			return
		}

		err = TopicApp.DeleteTopicByID(uint(id))
		if err != nil {
			respond.Error(w, http.StatusConflict, fmt.Errorf("%s", err.Error()))
			return
		}

		respond.Success(w, http.StatusOK, "Success")
	})
}
