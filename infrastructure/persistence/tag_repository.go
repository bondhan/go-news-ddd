package persistence

import (
	"github.com/bondhan/godddnews/domain"
	"github.com/bondhan/godddnews/domain/repository"
	"github.com/jinzhu/gorm"
)

type tagRepository struct {
	db *gorm.DB
}

//NewTagRepository ...
func NewTagRepository(newDB *gorm.DB) repository.TagRepository {
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
