package domain

import (
	"database/sql"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/guregu/null"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
)

// Topic ...
type Topic struct {
	gorm.Model
	Name    string `gorm:"column:name" json:"name" validate:"required"`
	Slug    string `gorm:"column:slug" json:"slug" validate:"required"`
	Version uint   `gorm:"column:version" json:"version"`
}

// TableName sets the insert table name for this struct type
func (m *Topic) TableName() string {
	return "m_topic"
}
