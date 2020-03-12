package repository

import (
	"github.com/bondhan/godddnews/domain"
	"github.com/jinzhu/gorm"
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

type topicRepository struct {
	db *gorm.DB
}

//NewTopicRepository ...
func NewTopicRepository(newDB *gorm.DB) TopicRepository {
	return &topicRepository{
		db: newDB,
	}
}

func (n *topicRepository) GetATopic(id uint) (domain.Topic, error) {
	var Topic domain.Topic
	err := n.db.Where("id = ?", id).First(&Topic).Error

	return Topic, err
}

func (n *topicRepository) InsertATopic(Topic domain.Topic) error {
	err := n.db.Create(&Topic).Error

	return err
}

func (n *topicRepository) UpdateTopic(Topic domain.Topic) error {
	ver := Topic.Version
	Topic.Version = ver + 1
	err := n.db.Model(&Topic).Where("version = ?", ver).UpdateColumns(Topic).Error //optimistic lock

	return err
}

func (n *topicRepository) GetATopicBySlug(slug string) (domain.Topic, error) {
	var topic domain.Topic
	err := n.db.Where("slug = ?", slug).First(&topic).Error

	return topic, err
}

func (n *topicRepository) GetATopicByID(id uint) (domain.Topic, error) {
	var topic domain.Topic
	err := n.db.Where("id = ?", id).First(&topic).Error

	return topic, err
}

func (n *topicRepository) GetAllTopics() ([]domain.Topic, error) {

	var nx []domain.Topic
	err := n.db.Find(&nx).Error

	return nx, err
}

func (n *topicRepository) DeleteTopic(topic domain.Topic) error {
	err := n.db.Unscoped().Delete(&topic).Error

	return err
}
