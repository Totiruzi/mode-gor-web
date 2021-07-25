package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", Home) 

	http.HandleFunc("/about", About)

	http.ListenAndServe(portNumber, nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
		b, err := fmt.Println("Go Forever love")
		if err !=nil {
			http.Error(w, "Could not process request", http.StatusBadRequest)
			return
		}
		fmt.Printf("The byte of the response is %d", b)
}

func About(w http.ResponseWriter, r *http.Request) {}

func renderTemplate(w http.ResponseWriter, temp string) {
	parseTemplate, _ := template.ParseFiles("./templates" + temp) 
	err := parseTemplate.Execute(w, nil)
	if err != nil {
		http.Error(w, "Problem parsing template", http.StatusBadRequest)
		return
	}
}