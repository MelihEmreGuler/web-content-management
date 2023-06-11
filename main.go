package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello World")

	http.HandleFunc("/home", Home)
	http.HandleFunc("/about", About)

	err := http.ListenAndServe(":8080", nil) // nil means use the default serve mux
	if err != nil {
		fmt.Println(err)
	}
}
