package main

import (
	"fmt"
	"github.com/MelihEmreGuler/web-content-management/pkg/handlers"
	"net/http"
)

func main() {
	fmt.Println("Hello World")

	http.HandleFunc("/home", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	err := http.ListenAndServe(":8080", nil) // nil means use the default serve mux
	if err != nil {
		fmt.Println(err)
	}
}
