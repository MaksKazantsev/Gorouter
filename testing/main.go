package main

import (
	gorouter "github.com/MaksKazantsev/Gorouter"
)

func main() {
	app := gorouter.NewApp(gorouter.WithAddress("3005"))

	app.GET("/test/{id}/{testing}", func(ctx *gorouter.Ctx) {
		id := ctx.Vars["testing"]
		ctx.Response.Write([]byte(id))
	})

	app.Listen()
}
