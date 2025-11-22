package Container

import (
	"encoding/gob"
	"errors"
	"fmt"
	"reflect"
)

type Container struct {
	pServices map[string]*Service
}

func New(services ...*Service) *Container {
	container := Container{
		pServices: make(map[string]*Service),
	}

	for _, service := range services {
		gob.Register(service)
	}

	return &container
}

func (container *Container) Register(service Service) *Container {
	name := getServiceName(service)

	if ser := container.pServices[name]; ser == nil {
		container.pServices[name] = &service
	}

	return container
}

func getServiceName(service Service) string {
	t := reflect.TypeOf(service)
	return t.Elem().Name()
}

func (container *Container) Get(service Service) (*Service, error) {
	name := getServiceName(service)

	val, ok := container.pServices[name]

	if !ok {
		return nil, errors.New(fmt.Sprintf("Service %s not found", name))
	}

	return val, nil
}
