package usecase

import (
	"github.com/bondhan/godddnews/domain/repository"
	"github.com/bondhan/godddnews/usecase/view"
	"github.com/sirupsen/logrus"
	"net/url"
)

type NewsApp interface {
	GetAllNews(queryStr url.Values) (view.NewsView, error)
	AddNews(newsData view.NewsData) error
}

type newsApp struct {
	newsRepo  repository.NewsRepository
	topicRepo repository.TopicRepository
	tagRepo   repository.TagRepository
}

func NewNewsApp(newsRepo repository.NewsRepository, topicRepo repository.TopicRepository, tagRepo repository.TagRepository) NewsApp {
	return &newsApp{
		newsRepo:  newsRepo,
		topicRepo: topicRepo,
		tagRepo:   tagRepo,
	}
}

// GetAllNews will get all news. Pagination and filter is also supported.
func (n *newsApp) GetAllNews(queryStr url.Values) (view.NewsView, error) {
	var newsView view.NewsView
	var err error

	logrus.Debug("GetAllNews")
	newsView, err = n.newsRepo.GetAllNews(queryStr)

	return newsView, err
}

// AddNews .
func (n *newsApp) AddNews(newsData view.NewsData) error {
	var err error

	return err
}
