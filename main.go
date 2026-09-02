package main

import "github.com/anotherlusitano/goProject/internal/app"

func main() {
	app, err := app.NewApp()

	if err != nil {
		panic(err)
	}

	app.Logger.Println("Setup is done!")
}
