package main

import (
	"goreact/pkg/handlers"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)

	http.HandleFunc("/about", handlers.About)

	log.Printf("Starting server on port: %s", portNumber)

	http.ListenAndServe(portNumber, nil)
}
