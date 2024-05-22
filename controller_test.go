package gorouter

import (
	"fmt"
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
	time.Sleep(time.Millisecond * 10)

	app.GET("/testing", func(ctx *Ctx) {
		_, _ = ctx.Response.Write([]byte("RECEIVED"))
		fmt.Println("RECEIVED GET")
	})

	getReq, err := http.NewRequest(http.MethodGet, "http://127.0.0.1:3005/testing", nil)
	require.NoError(t, err)

	cl := http.Client{}
	res, err := cl.Do(getReq)
	require.NoError(t, err)

	require.Equal(t, http.StatusOK, res.StatusCode)
}
