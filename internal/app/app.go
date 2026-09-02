package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/anotherlusitano/goProject/internal/api"
)

type App struct {
	Logger         *log.Logger
	WorkoutHandler *api.WorkoutHandler
}

func NewApp() (*App, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// TODO: add store

	workoutHandler := api.NewWorkoutHandler()

	app := &App{
		Logger:         logger,
		WorkoutHandler: workoutHandler,
	}

	return app, nil
}

func (a *App) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Status available")
}
