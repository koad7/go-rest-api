package main

import (
	"fmt"
	"net/http"

	transportHTTP "github.com/koad7/go-rest-api/internal/transport/http"
)

// App - teh struct
type App struct{}

// Run - sets up the  application
func (app *App) Run() error {
	fmt.Println("Setting up our API")

	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set up a server")
		return nil
	}
	return nil
}
func main() {
	fmt.Println("Go REST API Course")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up our REST API")
		fmt.Println(err)
	}
}
