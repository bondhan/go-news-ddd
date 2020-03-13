package manager

import (
	"sync"

	"github.com/bondhan/godddnews/application"
	"github.com/bondhan/godddnews/config"
	"github.com/bondhan/godddnews/infrastructure/persistence"
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
	container.Provide(application.NewNewsApp)
	container.Provide(application.NewTopicApp)
	container.Provide(application.NewTagApp)
	container.Provide(persistence.NewTopicRepository)
	container.Provide(persistence.NewTagRepository)
	container.Provide(persistence.NewNewsRepository)

	return container
}

// GetContainer ...
func GetContainer() *dig.Container {
	once.Do(func() {

		singleton = &Manager{buildContainer()}
	})

	return singleton.container
}
