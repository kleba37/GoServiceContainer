package Container

import (
	"testing"

	"github.com/kleba37/GoServiceContainer"
	"go.uber.org/goleak"
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
		di := GoServiceContainer.New()
		myService := &TestService{Message: "Hello, World!"}
		di.Register(myService)

		s, err := di.Get(myService)

		if err != nil {
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

func TestGetIsNotAccessService(t *testing.T) {
	t.Run("Testing GetIsNotAccessService", func(t *testing.T) {
		di := GoServiceContainer.New()
		myService := &TestService{Message: "Hello, World!"}
		di.Register(myService)
		_, err := di.Get(myService)

		if err != nil {
			t.Fatalf("Должна была придти ошибка, но не пришла")
		}
	})
}

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

func TestGoRoutine(t *testing.T) {
	goleak.VerifyNone(t)
}
