package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/anotherlusitano/goProject/internal/api"
	"github.com/anotherlusitano/goProject/internal/store"
	"github.com/anotherlusitano/goProject/migrations"
)

type App struct {
	Logger         *log.Logger
	WorkoutHandler *api.WorkoutHandler
	DB             *sql.DB
}

func NewApp() (*App, error) {
	pgDb, err := store.Open()
	if err != nil {
		return nil, err
	}

	err = store.MigrateFS(pgDb, migrations.FS, ".")
	if err != nil {
		panic(err)
	}

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// TODO: add store

	workoutHandler := api.NewWorkoutHandler()

	app := &App{
		Logger:         logger,
		WorkoutHandler: workoutHandler,
		DB:             pgDb,
	}

	return app, nil
}

func (a *App) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Status available")
}
