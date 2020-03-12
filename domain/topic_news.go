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

// TopicNews ...
type TopicNews struct {
	NewsID  uint `gorm:"column:news_id" json:"news_id"`
	TopicID uint `gorm:"column:topic_id" json:"topic_id"`
}

// TableName sets the insert table name for this struct type
func (m *TopicNews) TableName() string {
	return "m_topic_news"
}
