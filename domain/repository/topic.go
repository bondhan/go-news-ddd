package repository

import (
	"database/sql"
	"github.com/bondhan/godddnews/domain"
)

// TopicRepository ...
type TopicRepository interface {
	InsertATopic(Topic domain.Topic) error
	GetAllTopics() ([]domain.Topic, error)
}

type topicRepository struct {
	db *sql.DB
}

//NewTopicRepository ...
func NewTopicRepository(newDB *sql.DB) TopicRepository {
	return &topicRepository{
		db: newDB,
	}
}

func (n *topicRepository) InsertATopic(Topic domain.Topic) error {

	return nil
}

func (n *topicRepository) GetAllTopics() ([]domain.Topic, error) {

	var nx []domain.Topic

	return nx, nil
}
