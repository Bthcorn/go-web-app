package main

import (
	"log"

	app "github.com/Bthchth/go_app/app"
)

func main() {
	a, err := app.NewApp()
	if err != nil {
		log.Fatal(err)
	}
	// create a router
	// r := http.NewServeMux()

	// register the handler
	a.RegisterHandler("/{$}", a.Index)
	a.RegisterHandler("/about", a.About)
	a.RegisterHandler("/users", a.Users)
	a.RegisterHandler("/users/{username}", a.User)

	// log.Printf("Server started on port 3000\n")
	// start the server
	a.Serve(3000)
}
