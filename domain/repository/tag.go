package repository

import (
	"github.com/bondhan/godddnews/domain"
	"github.com/jinzhu/gorm"
)

// TagRepository ...
type TagRepository interface {
	GetATag(id uint) (domain.Tag, error)
	GetATagBySlug(slug string) (domain.Tag, error)
	InsertATag(Tag domain.Tag) error
	GetAllTags() ([]domain.Tag, error)
	UpdateTag(Tag domain.Tag) error
	DeleteTag(tag domain.Tag) error
	GetATagByID(id uint) (domain.Tag, error)
}

type tagRepository struct {
	db *gorm.DB
}

//NewTagRepository ...
func NewTagRepository(newDB *gorm.DB) TagRepository {
	return &tagRepository{
		db: newDB,
	}
}

func (n *tagRepository) GetATag(id uint) (domain.Tag, error) {
	var Tag domain.Tag
	err := n.db.Where("id = ?", id).First(&Tag).Error

	return Tag, err
}

func (n *tagRepository) InsertATag(Tag domain.Tag) error {
	err := n.db.Create(&Tag).Error

	return err
}

func (n *tagRepository) UpdateTag(Tag domain.Tag) error {
	ver := Tag.Version
	Tag.Version = ver + 1
	err := n.db.Model(&Tag).Where("version = ?", ver).UpdateColumns(Tag).Error //optimistic lock

	return err
}

func (n *tagRepository) GetATagBySlug(slug string) (domain.Tag, error) {
	var tag domain.Tag
	err := n.db.Where("slug = ?", slug).First(&tag).Error

	return tag, err
}

func (n *tagRepository) GetATagByID(id uint) (domain.Tag, error) {
	var tag domain.Tag
	err := n.db.Where("id = ?", id).First(&tag).Error

	return tag, err
}

func (n *tagRepository) GetAllTags() ([]domain.Tag, error) {

	var nx []domain.Tag
	err := n.db.Find(&nx).Error

	return nx, err
}

func (n *tagRepository) DeleteTag(tag domain.Tag) error {
	err := n.db.Unscoped().Delete(&tag).Error

	return err
}
