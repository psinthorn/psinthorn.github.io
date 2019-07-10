package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello, This is my 1st personal webapp by Golang")

	//-----------------------------------------------------------------------------
	// #GO http
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, This is my 1st personal webapp by Golang")
	})

	http.HandleFunc("/about-go", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "About Me")
	})

	log.Fatal(http.ListenAndServe(":8888", nil))

	//-------------------------------------------------------------------------------
}
