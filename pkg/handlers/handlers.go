package handlers

import (
	"net/http"
	"goreact/pkg/rendler"
)

func Home(w http.ResponseWriter, r *http.Request) {
	randler.RenderTemplate(w, "home.page.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	randler.RenderTemplate(w, "about.page.html")
}
