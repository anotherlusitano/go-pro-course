package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/anotherlusitano/goProject/internal/app"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "go backend server port")
	flag.Parse()

	app, err := app.NewApp()

	if err != nil {
		panic(err)
	}

	app.Logger.Println("Setup is done!")

	http.HandleFunc("/health", HealthCheck)

	// Server config
	var (
		address      = fmt.Sprintf(":%d", port)
		idleTimeout  = time.Minute
		readTimeout  = 10 * time.Second
		writeTimeout = 30 * time.Second
	)

	server := &http.Server{
		Addr:         address,
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

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Status available")
}
