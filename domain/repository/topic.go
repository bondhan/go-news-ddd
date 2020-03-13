package repository

import (
	"github.com/bondhan/godddnews/domain"
)

// TopicRepository ...
type TopicRepository interface {
	GetATopic(id uint) (domain.Topic, error)
	GetATopicBySlug(slug string) (domain.Topic, error)
	InsertATopic(Topic domain.Topic) error
	GetAllTopics() ([]domain.Topic, error)
	UpdateTopic(Topic domain.Topic) error
	DeleteTopic(topic domain.Topic) error
	GetATopicByID(id uint) (domain.Topic, error)
}
