package main

import (
	// "goreact/pkg/config"
	"goreact/pkg/config"
	"goreact/pkg/handlers"
	"goreact/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Can not create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)

	http.HandleFunc("/about", handlers.Repo.About)

	log.Printf("Starting server on port: %s", portNumber)

	http.ListenAndServe(portNumber, nil)
}
