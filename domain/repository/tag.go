package repository

import (
	"github.com/bondhan/godddnews/domain"
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
