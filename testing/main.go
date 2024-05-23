package main

import (
	gorouter "github.com/MaksKazantsev/Gorouter"
)

func main() {
	app := gorouter.NewApp(gorouter.WithAddress("3002"))

	app.GET("/test/{id}", func(ctx *gorouter.Ctx) {
		
		ctx.Response.Write([]byte("hello it works"))
	})

	app.Listen()
}
