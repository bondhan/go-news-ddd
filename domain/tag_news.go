package domain

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
)

// TagNews ...
type TagNews struct {
	NewsID uint `gorm:"column:news_id" json:"news_id"`
	TagID  uint `gorm:"column:tag_id" json:"tag_id"`
}

// TableName sets the insert table name for this struct type
func (m *TagNews) TableName() string {
	return "m_tag_news"
}
