package client

import (
	"fmt"
)

type manager struct {
	objects map[string]interface{}
}

type Manager interface {
	SetObject(name string, object interface{})
	GetObject(name string) interface{}
}

func NewManager() Manager {
	obj := make(map[string]interface{})
	return &manager{
		objects: obj,
	}
}

// SetObject sets an object to container by name
func (m *manager) SetObject(name string, object interface{}) {
	m.objects[name] = object
}

// GetObject returns object from container using specified name.
// Need to cast to a concrete type.
func (m *manager) GetObject(name string) interface{} {
	return m.objects[createKey(name, "")]
}

func createKey(name string, ID string) string {
	key := name
	if ID != "" {
		key = fmt.Sprintf("%s%s%s", name, "#", ID)
	}
	return key
}
