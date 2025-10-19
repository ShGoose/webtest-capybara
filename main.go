package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func secondHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is the second handler!")
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/second", secondHandler)
	http.ListenAndServe(":8080", nil)
}
