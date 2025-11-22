package main

import (
	"fmt"
	"log"
	_ "testing"

	"github.com/kleba37/GoServiceContainer"
)

type Service struct {
	Message string
}

func (service *Service) SayHello() string {
	return service.Message
}

func main() {
	di := GoServiceContainer.New()
	mySer := &Service{
		Message: "Hello, World!",
	}

	di.Register(mySer)

	instance := &Service{}

	s, err := di.Get(instance)

	if err != nil {
		log.Fatalf("Don't get service error")
	}

	service, ok := (*s).(*Service)

	if !ok {
		log.Fatal(fmt.Sprintf("Don't casting %s", service))
	}

	log.Println("Get service message", service.SayHello())
}
