package utils

import (
	"errors"
	"regexp"
	"sync"

	v "gopkg.in/go-playground/validator.v9"
)

var (
	correctSlug *regexp.Regexp
	once        sync.Once
)

// ValidateModels ..
func ValidateModels(mod interface{}) error {
	err := v.New().Struct(mod)

	return err
}

func ValidateSlug(slug string) error {
	once.Do(func() {
		correctSlug = regexp.MustCompile(`^[a-z-]+$`)
	})

	if !correctSlug.MatchString(slug) {
		return errors.New("slug must be lower case letters with - (dash) only")
	}

	return nil
}
