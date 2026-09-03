package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/anotherlusitano/goProject/internal/app"
	"github.com/anotherlusitano/goProject/internal/routes"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "go backend server port")
	flag.Parse()

	app, err := app.NewApp()

	if err != nil {
		panic(err)
	}

	defer app.DB.Close()

	app.Logger.Println("Setup is done!")

	// Server config
	var (
		address      = fmt.Sprintf(":%d", port)
		r            = routes.SetupRoutes(app)
		idleTimeout  = time.Minute
		readTimeout  = 10 * time.Second
		writeTimeout = 30 * time.Second
	)

	server := &http.Server{
		Addr:         address,
		Handler:      r,
		IdleTimeout:  idleTimeout,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	app.Logger.Printf("We are running on port: %d\n", port)

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}
