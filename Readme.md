# Go Service Container with DI

## Init DI for App
```go 
di := (&GoServiceContainer.DI{}).New()
myService := &TestService{Message: "Hello, World!"}
di.Register(myService)
```

## Get service from DI

```go
di.Get(myService)
```

## Autorun services example
```go
var autorun = map[string]Container.Service{
    "DB": db(),
}

...

func db() *sql.DB {
    err := godotenv.Load()
    
    if err != nil {
        return nil
    }
    
    dsn := os.Getenv("DB_DSN")
    
    if len(dsn) == 0 {
        panic("DSN invalid")
    }
    
    pool, err := sql.Open(os.Getenv("DB_CONNECTION"), dsn)
    
    if err != nil {
        fmt.Println("DB error connection")
        panic(err)
    }
    
    log.Println("DB connect with DSN: ", dsn)
    
    return pool
}
```

