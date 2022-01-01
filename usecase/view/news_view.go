package view

import (
	"github.com/bondhan/godddnews/domain"
	_ "gopkg.in/go-playground/validator.v9" //need to validate the models
)

// NewsView ...
type NewsView struct {
	Total      int           `json:"total"`
	TotalPage  int           `json:"total_page"`
	PageNumber int           `json:"page_number"`
	PageSize   int           `json:"page_size"`
	News       []domain.News `json:"news"`
}

type NewsData struct {
	ID         uint     `json:"id" validate:"omitempty"`
	Title      string   `json:"title" validate:"required"`
	Slug       string   `json:"slug" validate:"required"`
	Content    string   `json:"content"`
	Status     string   `json:"status" validate:"omitempty,oneof=draft publish deleted"`
	TopicSlugs []string `json:"topic_slugs" validate:"required"`
	TagSlugs   []string `json:"tag_slugs"`
}
