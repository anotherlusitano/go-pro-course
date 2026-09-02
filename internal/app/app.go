package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type App struct {
	Logger *log.Logger
}

func NewApp() (*App, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &App{
		Logger: logger,
	}

	return app, nil
}

func (a *App) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Status available")
}
