package gorouter

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"
)

var app *App

func start() {
	app = NewApp(WithAddress("3005"))
	app.Listen()
}

func TestGet(t *testing.T) {
	go start()
	time.Sleep(time.Second)

	app.GET("/test", func(ctx *Ctx) {
		_, _ = ctx.Response.Write([]byte("received"))
	})

	req, err := http.NewRequest(http.MethodGet, "http://127.0.0.1:3005/test", nil)
	require.NoError(t, err)

	client := http.Client{}
	res, err := client.Do(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, res.StatusCode)
}
