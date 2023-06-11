package main

import (
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "This is the home page")
	renderTemplate(w, "home.page.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "This is the about page")
	renderTemplate(w, "about.page.html")
}
