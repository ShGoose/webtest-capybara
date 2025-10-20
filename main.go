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

func sum(a, b int) string {
	result := fmt.Sprint(a + b)
	return fmt.Sprint("Result is: %n", result)
}

func thirdHandler(w http.ResponseWriter, r *http.Request) {
	message := sum(10, 11)
	fmt.Fprintln(w, message)
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/second", secondHandler)
	http.HandleFunc("/third", thirdHandler)
	http.ListenAndServe(":8080", nil)
}
