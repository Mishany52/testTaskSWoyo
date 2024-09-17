package application

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Mishany52/testTaskSWoyo/config"
)

type App struct {
	router http.Handler
}

func New() *App {
	app := &App{
		router: loadRoutes(),
	}
	return app
}

func (a *App) Start(ctx context.Context) error {
	var conf = config.New()
	server := &http.Server{
		Addr: conf.ServerAddr,
		Handler: a.router,
	}

	err := server.ListenAndServe()
	if err!= nil{
		return fmt.Errorf("failed to listen to server: %w", err)
	}

	return nil
}