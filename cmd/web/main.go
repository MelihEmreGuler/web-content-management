package main

import (
	"fmt"
	"github.com/MelihEmreGuler/web-content-management/pkg/config"
	"github.com/MelihEmreGuler/web-content-management/pkg/handlers"
	"github.com/MelihEmreGuler/web-content-management/pkg/render"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"
)

var app config.AppConfig        // create a variable to hold the app config
var session *scs.SessionManager // create a variable to hold the session manager

func main() {
	fmt.Println("Starting application on port 8080")

	app.InProduction = false // set to true in production

	session = scs.New()                            // create a new session manager
	session.Lifetime = 24 * time.Hour              // set the lifetime of the session to 24 hours
	session.Cookie.Persist = true                  // persist the session across browser restarts
	session.Cookie.SameSite = http.SameSiteLaxMode // set the SameSite cookie mode to Lax
	session.Cookie.Secure = app.InProduction       // set the Secure cookie mode to false for development

	app.Session = session // assign the session manager to the app config

	tc, err := render.CreateTemplateCache() // create a template cache
	if err != nil {
		fmt.Println("Cannot create template cache")
	}

	app.TemplateCache = tc // assign the template cache to the app config

	// set to true in production mode to use the template cache,
	//set false in development mode to create a new template cache every time
	app.UseCache = true

	repo := handlers.NewRepo(&app) // create a new repository
	handlers.NewHandlers(repo)     // create new handlers and assign the repository
	render.NewTemplates(&app)      // pass the app config to the render package

	srv := &http.Server{
		Addr:    ":8080",
		Handler: routes(),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
