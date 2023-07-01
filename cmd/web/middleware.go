package main

import (
	"fmt"
	"github.com/justinas/nosurf"
	"net/http"
)

// WriteToConsole writes to the console the path and method of the request
func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page \"" + r.URL.Path + "\"" + " with method \"" + r.Method + "\"" + " from IP address \"" + r.RemoteAddr + "\"")
		next.ServeHTTP(w, r)
	})
}

// NoSurf adds CSRF protection to all POST requests
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{ // set the base cookie
		HttpOnly: true,                 // cookie cannot be accessed by JavaScript
		Path:     "/",                  // cookie is valid for all pages
		Secure:   app.InProduction,     // cookie will only be sent over HTTPS
		SameSite: http.SameSiteLaxMode, // cookie will not be sent on cross-site requests
	})
	return csrfHandler
}

// SessionLoad loads and saves the session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
