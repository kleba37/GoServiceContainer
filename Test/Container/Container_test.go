package Container

import (
	"testing"

	"GoServiceContainer/GoServiceContainer"
)

type TestServiceInterface interface {
	SayHello() string
}

type TestService struct {
	Message string
}

func (s *TestService) SayHello() string {
	return s.Message
}

func TestRegisterAndGetConcreteService(t *testing.T) {
	t.Run("Testing register service", func(t *testing.T) {
		di := (&GoServiceContainer.DI{}).New()
		myService := &TestService{Message: "Hello, World!"}
		di.Register(myService)

		s := di.Get(myService)

		if s == nil {
			t.Fatal("Get() вернул nil, сервис не найден в контейнере")
		}

		service, ok := (*s).(TestServiceInterface)
		if !ok {
			t.Fatalf("Не удалось преобразовать полученный сервис к типу TestServiceInterface. Получен тип: %T", s)
		}

		expectedMessage := "Hello, World!"
		if message := service.SayHello(); message != expectedMessage {
			t.Errorf("Ожидалось сообщение '%s', но получено '%s'", expectedMessage, message)
		}
	})

}
