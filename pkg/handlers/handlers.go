package handlers

import (
	"github.com/MelihEmreGuler/web-content-management/pkg/render"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "This is the home page")
	render.RenderTemplate(w, "home.page.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "This is the about page")
	render.RenderTemplate(w, "about.page.html")
}
