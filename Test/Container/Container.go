package Container

import (
	"github.com/kleba/Container/internal/DI"
	"testing"
)

// TestServiceInterface определяет интерфейс для нашего тестового сервиса.
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
	di := (&DI.DI{}).New()

	myService := TestService{Message: "Hello, World!"}

	di.Register(myService)

	retrievedServicePtr := di.Get(myService)

	if retrievedServicePtr == nil {
		t.Fatal("Get() вернул nil, сервис не найден в контейнере")
	}

	retrievedService := retrievedServicePtr

	service, ok := (*retrievedService).(TestServiceInterface)
	if !ok {
		t.Fatalf("Не удалось преобразовать полученный сервис к типу TestServiceInterface. Получен тип: %T", retrievedService)
	}

	expectedMessage := "Hello, World!"
	if message := service.SayHello(); message != expectedMessage {
		t.Errorf("Ожидалось сообщение '%s', но получено '%s'", expectedMessage, message)
	}
}