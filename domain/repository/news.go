package repository

import (
	"github.com/bondhan/godddnews/domain"
	"github.com/bondhan/godddnews/interfaces/handlers/database"
	"github.com/bondhan/godddnews/usecase/view"
	"github.com/jinzhu/gorm"
	"net/url"
)

const queryPageSizeDefault = 10
const queryPageNumberDefault = 1

// NewsRepository ...
type NewsRepository interface {
	GetAllNews(queryStr url.Values) (view.NewsView, error)
	InsertNews(News domain.News, NewTopics []domain.Topic, NewTags []domain.Tag) error
}

type newsRepository struct {
	db        database.SQLHandler
	topicRepo TopicRepository
	tagRepo   TagRepository
}

//NewNewsRepository ...
func NewNewsRepository(newDB *gorm.DB, topicRepo TopicRepository, tagRepo TagRepository) NewsRepository {
	return &newsRepository{
		db:        newDB,
		topicRepo: topicRepo,
		tagRepo:   tagRepo,
	}
}

func (n *newsRepository) GetAllNews(queryStr url.Values) (view.NewsView, error) {
	return view.NewsView{}, nil
}

func (n *newsRepository) InsertNews(News domain.News, NewTopics []domain.Topic, NewTags []domain.Tag) error {
	return nil
}
