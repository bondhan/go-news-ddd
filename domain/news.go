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

// News ...
type News struct {
	gorm.Model
	Title   string  `gorm:"column:title" json:"title"`
	Slug    string  `gorm:"column:slug" json:"slug"`
	Content string  `gorm:"column:content" json:"content"`
	Status  string  `gorm:"column:status" json:"status"`
	Version uint    `gorm:"column:version" json:"version"`
	Topics  []Topic `gorm:"many2many:m_topic_news;"`
	Tags    []Tag   `gorm:"many2many:m_tag_news;"`
}

// TableName sets the insert table name for this struct type
func (m *News) TableName() string {
	return "m_news"
}
