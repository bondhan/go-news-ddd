package usecase

import (
	"errors"
	"github.com/bondhan/godddnews/infrastructure/utils"

	"github.com/bondhan/godddnews/domain"
	"github.com/bondhan/godddnews/domain/repository"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type TopicApp interface {
	GetAllTopics() ([]domain.Topic, error)
	AddTopic(topic domain.Topic) error
	GetTopicBySlug(slug string) (domain.Topic, error)
	GetTopicByID(id uint) (domain.Topic, error)
	UpdateTopicBySlug(topic domain.Topic, slug string) error
	DeleteTopicBySlug(slug string) error
	UpdateTopicByID(topic domain.Topic, id uint) error
	DeleteTopicByID(ID uint) error
}

type topicApp struct {
	topicRepo repository.TopicRepository
}

func NewTopicApp(topicRepo repository.TopicRepository) TopicApp {
	return &topicApp{
		topicRepo: topicRepo,
	}
}

// GetAllTopics ...
func (n *topicApp) GetAllTopics() ([]domain.Topic, error) {
	var err error

	logrus.Debug("GetAllTopics")
	topics, err := n.topicRepo.GetAllTopics()

	return topics, err
}

func (n *topicApp) AddTopic(topic domain.Topic) error {
	var err error

	logrus.Debug("AddTopic")

	err = utils.ValidateModels(topic)
	if err != nil {
		return err
	}

	err = utils.ValidateSlug(topic.Slug)
	if err != nil {
		return err
	}

	tp, err := n.topicRepo.GetATopicBySlug(topic.Slug)
	//if error was not caused by empty select, means something happened during query to db
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return err
	}

	//if returned non empty data
	if (tp.Model != gorm.Model{}) {
		return errors.New("Data already exist")
	}

	if topic.Version < 1 {
		topic.Version = 1
	}

	err = n.topicRepo.InsertATopic(topic)

	return err
}

// GetTopicBySlug ...
func (n *topicApp) GetTopicBySlug(slug string) (domain.Topic, error) {

	logrus.Debug("GetTopicBySlug")
	topic, err := n.topicRepo.GetATopicBySlug(slug)
	if gorm.IsRecordNotFoundError(err) {
		err = errors.New("Data not found")
	}

	return topic, err
}

func (n *topicApp) GetTopicByID(id uint) (domain.Topic, error) {

	logrus.Debug("GetTopicByID")
	topic, err := n.topicRepo.GetATopicByID(id)
	if gorm.IsRecordNotFoundError(err) {
		err = errors.New("Data not found")
	}

	return topic, err
}

func (n *topicApp) UpdateTopicBySlug(topic domain.Topic, oldslug string) error {

	var err error

	logrus.Debug("UpdateTopicBySlug")

	err = utils.ValidateModels(topic)
	if err != nil {
		return err
	}

	err = utils.ValidateSlug(topic.Slug)
	if err != nil {
		return err
	}

	tp, err := n.topicRepo.GetATopicBySlug(oldslug)
	//if error was not caused by empty select, means something happened during query to db
	if gorm.IsRecordNotFoundError(err) {
		return errors.New("Data to be updated not exist")
	}

	//if returned non empty data
	if (tp.Model == gorm.Model{}) {
		return errors.New("Data to be updated not exist")
	}

	tp.Name = topic.Name
	tp.Slug = topic.Slug

	err = n.topicRepo.UpdateTopic(tp)

	return err
}

func (n *topicApp) UpdateTopicByID(topic domain.Topic, id uint) error {

	var err error

	logrus.Debug("UpdateTopicByID")

	err = utils.ValidateModels(topic)
	if err != nil {
		return err
	}

	err = utils.ValidateSlug(topic.Slug)
	if err != nil {
		return err
	}

	tp, err := n.topicRepo.GetATopicByID(id)
	//if error was not caused by empty select, means something happened during query to db
	if gorm.IsRecordNotFoundError(err) {
		return errors.New("Data to be updated not exist")
	}

	//if returned non empty data
	if (tp.Model == gorm.Model{}) {
		return errors.New("Data to be updated not exist")
	}

	tp.Name = topic.Name
	tp.Slug = topic.Slug

	err = n.topicRepo.UpdateTopic(tp)

	return err
}

func (n *topicApp) DeleteTopicBySlug(slug string) error {

	var err error

	logrus.Debug("DeleteTopicBySlug")

	tp, err := n.topicRepo.GetATopicBySlug(slug)
	//if error was not caused by empty select, means something happened during query to db
	if gorm.IsRecordNotFoundError(err) {
		return errors.New("Data to be deleted not exist")
	}

	//if returned non empty data
	if (tp.Model == gorm.Model{}) {
		return errors.New("Data to be deleted not exist")
	}

	err = n.topicRepo.DeleteTopic(tp)
	if err != nil {
		return errors.New("Resource is still being used or DB operation error")
	}

	return err
}

func (n *topicApp) DeleteTopicByID(id uint) error {

	var err error

	logrus.Debug("DeleteTopicByID")

	tp, err := n.topicRepo.GetATopicByID(id)
	//if error was not caused by empty select, means something happened during query to db
	if gorm.IsRecordNotFoundError(err) {
		return errors.New("Data to be deleted not exist")
	}

	//if returned non empty data
	if (tp.Model == gorm.Model{}) {
		return errors.New("Data to be deleted not exist")
	}

	err = n.topicRepo.DeleteTopic(tp)
	if err != nil {
		return errors.New("Resource is still being used or DB operation error")
	}

	return err
}
