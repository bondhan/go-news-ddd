package requests

import (
	"encoding/json"
	. "github.com/go-ozzo/ozzo-validation"
	"net/http"
)

type news struct {
	Title      string   `json:"title"`
	Slug       string   `json:"slug"`
	Content    string   `json:"content"`
	Status     string   `json:"status"`
	TopicSlugs []string `json:"topicSlugs"`
	TagSlugs   []string `json:"tagSlugs"`
}

func (n news) validator() error {
	return ValidateStruct(&n,
		Field(&n.Title, Required),
		Field(&n.Content, Required),
		Field(&n.Slug, Required),
		Field(&n.Status, Required, In("draft", "deleted", "publish")),
		Field(&n.TopicSlugs, Required),
		Field(&n.TagSlugs, Required),
	)
}

func NewNewsRequest(r http.Request) (news, error) {
	var n news

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&n); err != nil {
		return n, err
	}

	return n, nil
}
