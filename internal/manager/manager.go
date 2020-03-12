package manager

import (
	"sync"

	"github.com/bondhan/godddnews/application"
	"github.com/bondhan/godddnews/config"
	"github.com/bondhan/godddnews/domain/repository"
	"go.uber.org/dig"
)

// Manager ...
type Manager struct {
	container *dig.Container
}

var (
	singleton *Manager
	once      sync.Once
)

// BuildContainer ...
func buildContainer() *dig.Container {
	container := dig.New()
	container.Provide(config.NewDbConfig)
	container.Provide(repository.NewNewsRepository)
	container.Provide(repository.NewTopicRepository)
	container.Provide(repository.NewTagRepository)
	container.Provide(application.NewNewsApp)
	container.Provide(application.NewTopicApp)
	container.Provide(application.NewTagApp)

	return container
}

// GetContainer ...
func GetContainer() *dig.Container {
	once.Do(func() {

		singleton = &Manager{buildContainer()}
	})

	return singleton.container
}
