package main

import "fmt"

// App - teh struct 
type App struct{}

// Run - sets up the  application 
func (app *App) Run() error {
	fmt.Println("Setting up our API")
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
