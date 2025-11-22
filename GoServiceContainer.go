package GoServiceContainer

import (
	"github.com/kleba37/GoServiceContainer/pkg/Container"
	//_ "modernc.org/sqlite"
)

var autorun = map[string]Container.Service{
	//"DB": db(),
}

type DI struct{}

func New() *Container.Container {
	container := Container.New()

	for _, ser := range autorun {
		container.Register(ser)
	}

	return container
}

// Пример autoload
//func db() *sql.DB {
//	err := godotenv.Load()
//	if err != nil {
//		return nil
//	}
//
//	dsn := os.Getenv("DB_DSN")
//
//	if len(dsn) == 0 {
//		panic("DSN invalid")
//	}
//
//	pool, err := sql.Open(os.Getenv("DB_CONNECTION"), dsn)
//
//	if err != nil {
//		fmt.Println("DB error connection")
//		panic(err)
//	}
//		fmt.Println("DB connect with DSN: ", dsn)
//	return pool
//}
