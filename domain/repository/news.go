package repository

import (
	"net/url"

	"github.com/bondhan/godddnews/application/view"
	"github.com/bondhan/godddnews/domain"
)

// NewsRepository ...
type NewsRepository interface {
	GetAllNews(queryStr url.Values) (view.NewsView, error)
	InsertNews(News domain.News, NewTopics []domain.Topic, NewTags []domain.Tag) error
	GetNewsBySlug(slug string) (domain.News, error)
	GetNewsByID(ID uint) (domain.News, error)
	UpdateNews(News domain.News, NewTopics []domain.Topic, NewTags []domain.Tag) error
	DeleteNewsBySlug(News domain.News, OldTopics []domain.Topic, OldTags []domain.Tag) error
	GetNewsByTopicSlug(slug string) ([]*domain.News, error)
	GetNewsByTagSlug(slug string) ([]*domain.News, error)
}
