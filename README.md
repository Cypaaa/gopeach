# 🍑 GoPeach

**GoPeach** is a [fasthttp](github.com/valyala/fasthttp) based framework 😃 made to simplify [fasthttp](github.com/valyala/fasthttp) while keeping its performance.  
**GoPeach** also integrats a database structure to allow you to implement gorm easily.

```bash
go get -u github.com/Cypaaa/gopeach
```

## Quickstart

```go
package main

import (
    "github.com/Cypaaa/gopeach"
)

func main() {
    app := gopeach.New()

    app.Get("/", func(ctx *gopeach.RequestCtx) {
      ctx.Send("Hello World!")
    })

    app.Listen(":8080")
}
```

## 🔒 JWT and Scrypt

**GoPeach** implements 4 functions to help you use JWT (EdDSA) and Scrypt.  

- JWT: GenerateJWT, VerifyJWT 
- Scrypt: Hash, CompareHash

You will find examples in the example directory.

## 👨‍💻👩‍💻 Middlewares

You can add custom or community made middlewares!  
They handle a request in order they are defined.  
Your first middleware will be executed before your second one which itself will execute before your handler

```go
func main() {
    app := gopeach.New()

    // Your first middleware
    app.Middleware(func(ctx *gopeach.RequestCtx) {
        fmt.Println("This is your first middleware")

        // Pass to the next handler only if method is "GET"
        // else the request won't be handled anymore
        if ctx.MethodString() == "GET" {
            ctx.Next()
            return
        }
    })

    // Your second middleware
    app.Middleware(func(ctx *gopeach.RequestCtx) {
        fmt.Println("This is your second one")
        ctx.Next()
    })

    // Get all routes
    app.Get("/", func(ctx *gopeach.RequestCtx) {
        ctx.Send("Hello World!")
    })

    app.Listen(":8080")
}
```

## 📓 Database

```go
func main() {
    godotenv.Load()

    // Dialector
    dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_NAME"),
    )
    dialector := mysql.Open(dns)
    
    // Database
    db, err := gopeach.NewDatabase(dialector)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    err = db.Migrate(
        models.MyModelTest{},
        // ...
    )
    if err != nil {
        log.Fatal(err)
    }
}
```
