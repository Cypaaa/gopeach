package main

import (
	"log"

	"github.com/Cypaaa/gopeach"
)

func main() {
	app := gopeach.New()

	// Detects routes in the order they are defined
	// Routes are case sensitive

	app.Get("/", func(ctx *gopeach.RequestCtx) {
		ctx.Send("Hello World!")
	})

	app.Get("/:id", func(ctx *gopeach.RequestCtx) {
		ctx.Send("ID is " + ctx.Params["id"])
	})

	// this route takes priority over the next one
	app.Get("/match/123", func(ctx *gopeach.RequestCtx) {
		ctx.Send("Match 123")
	})

	// so this route will match "/anything/123" or "/somethingelse/123 but not /match/123"
	app.Get("/:anything/123", func(ctx *gopeach.RequestCtx) {
		ctx.Send(ctx.Params["anything"] + " but not Match 123")
	})

	app.Get("/:even with space/123/", func(ctx *gopeach.RequestCtx) {
		ctx.Send(ctx.Params["anything"] + " but not Match 123")
	})

	// listen on port 8080 and fatal log if it fails
	log.Fatal(app.Listen(":8080"))
}
