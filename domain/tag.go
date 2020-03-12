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

// Tag ...
type Tag struct {
	gorm.Model
	Name    string `gorm:"column:name" json:"name"`
	Slug    string `gorm:"column:slug" json:"slug"`
	Version uint   `gorm:"column:version" json:"version"`
}

// TableName sets the insert table name for this struct type
func (m *Tag) TableName() string {
	return "m_tag"
}
