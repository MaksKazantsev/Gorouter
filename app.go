package gorouter

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type App struct {
	s *http.Server
	*Controller
}

// NewApp inits new app
func NewApp(o ...Options) *App {
	opt := OptionsStruct{
		Addr:   PORT_DEFAULT,
		WriteT: WRITE_DEFAULT,
		ReadT:  READ_DEFAULT,
	}

	for _, options := range o {
		options(&opt)
	}

	ctrl := NewController()
	router := NewRouter(ctrl)

	s := &http.Server{
		Addr:         opt.Addr,
		WriteTimeout: opt.WriteT,
		ReadTimeout:  opt.ReadT,
		Handler:      router,
	}
	return &App{
		s,
		ctrl,
	}
}

// Listen starts http server with provided options in app
func (a *App) Listen() {
	fmt.Printf("Server listen on port: %s", a.s.Addr)
	if err := http.ListenAndServe(a.s.Addr, nil); err != nil {
		panic("failed to listen: " + err.Error())
	}
}

// Shutdown make graceful shutdown
func (a *App) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := a.s.Shutdown(ctx); err != nil {
		panic("failed to shutdown app: " + err.Error())
	}
}
