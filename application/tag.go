package application

import (
	"errors"

	"github.com/bondhan/godddnews/domain"
	"github.com/bondhan/godddnews/domain/repository"
	"github.com/bondhan/godddnews/internal/utils"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type TagApp interface {
	GetAllTags() ([]domain.Tag, error)
	AddTag(tag domain.Tag) error
	GetTagBySlug(slug string) (domain.Tag, error)
	GetTagByID(id uint) (domain.Tag, error)
	UpdateTagBySlug(tag domain.Tag, slug string) error
	DeleteTagBySlug(slug string) error
	UpdateTagByID(tag domain.Tag, id uint) error
	DeleteTagByID(ID uint) error
}

type tagApp struct {
	tagRepo repository.TagRepository
}

func NewTagApp(tagRepo repository.TagRepository) TagApp {
	return &tagApp{
		tagRepo: tagRepo,
	}
}

// GetAllTags ...
func (n *tagApp) GetAllTags() ([]domain.Tag, error) {
	var err error

	logrus.Debug("GetAllTags")
	tags, err := n.tagRepo.GetAllTags()

	return tags, err
}

func (n *tagApp) AddTag(tag domain.Tag) error {
	var err error

	logrus.Debug("AddTag")

	err = utils.ValidateModels(tag)
	if err != nil {
		return err
	}

	err = utils.ValidateSlug(tag.Slug)
	if err != nil {
		return err
	}

	tp, err := n.tagRepo.GetATagBySlug(tag.Slug)
	//if error was not caused by empty select, means something happened during query to db
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return err
	}

	//if returned non empty data
	if (tp.Model != gorm.Model{}) {
		return errors.New("Data already exist")
	}

	if tag.Version < 1 {
		tag.Version = 1
	}

	err = n.tagRepo.InsertATag(tag)

	return err
}

// GetTagBySlug ...
func (n *tagApp) GetTagBySlug(slug string) (domain.Tag, error) {

	logrus.Debug("GetTagBySlug")
	tag, err := n.tagRepo.GetATagBySlug(slug)
	if gorm.IsRecordNotFoundError(err) {
		err = errors.New("Data not found")
	}

	return tag, err
}

func (n *tagApp) GetTagByID(id uint) (domain.Tag, error) {

	logrus.Debug("GetTagByID")
	tag, err := n.tagRepo.GetATagByID(id)
	if gorm.IsRecordNotFoundError(err) {
		err = errors.New("Data not found")
	}

	return tag, err
}

func (n *tagApp) UpdateTagBySlug(tag domain.Tag, oldslug string) error {

	var err error

	logrus.Debug("UpdateTagBySlug")

	err = utils.ValidateModels(tag)
	if err != nil {
		return err
	}

	err = utils.ValidateSlug(tag.Slug)
	if err != nil {
		return err
	}

	tp, err := n.tagRepo.GetATagBySlug(oldslug)
	//if error was not caused by empty select, means something happened during query to db
	if gorm.IsRecordNotFoundError(err) {
		return errors.New("Data to be updated not exist")
	}

	//if returned non empty data
	if (tp.Model == gorm.Model{}) {
		return errors.New("Data to be updated not exist")
	}

	tp.Name = tag.Name
	tp.Slug = tag.Slug

	err = n.tagRepo.UpdateTag(tp)

	return err
}

func (n *tagApp) UpdateTagByID(tag domain.Tag, id uint) error {

	var err error

	logrus.Debug("UpdateTagByID")

	err = utils.ValidateModels(tag)
	if err != nil {
		return err
	}

	err = utils.ValidateSlug(tag.Slug)
	if err != nil {
		return err
	}

	tp, err := n.tagRepo.GetATagByID(id)
	//if error was not caused by empty select, means something happened during query to db
	if gorm.IsRecordNotFoundError(err) {
		return errors.New("Data to be updated not exist")
	}

	//if returned non empty data
	if (tp.Model == gorm.Model{}) {
		return errors.New("Data to be updated not exist")
	}

	tp.Name = tag.Name
	tp.Slug = tag.Slug

	err = n.tagRepo.UpdateTag(tp)

	return err
}

func (n *tagApp) DeleteTagBySlug(slug string) error {

	var err error

	logrus.Debug("DeleteTagBySlug")

	tp, err := n.tagRepo.GetATagBySlug(slug)
	//if error was not caused by empty select, means something happened during query to db
	if gorm.IsRecordNotFoundError(err) {
		return errors.New("Data to be deleted not exist")
	}

	//if returned non empty data
	if (tp.Model == gorm.Model{}) {
		return errors.New("Data to be deleted not exist")
	}

	err = n.tagRepo.DeleteTag(tp)
	if err != nil {
		return errors.New("Resource is still being used or DB operation error")
	}

	return err
}

func (n *tagApp) DeleteTagByID(id uint) error {

	var err error

	logrus.Debug("DeleteTagByID")

	tp, err := n.tagRepo.GetATagByID(id)
	//if error was not caused by empty select, means something happened during query to db
	if gorm.IsRecordNotFoundError(err) {
		return errors.New("Data to be deleted not exist")
	}

	//if returned non empty data
	if (tp.Model == gorm.Model{}) {
		return errors.New("Data to be deleted not exist")
	}

	err = n.tagRepo.DeleteTag(tp)
	if err != nil {
		return errors.New("Resource is still being used or DB operation error")
	}

	return err
}
