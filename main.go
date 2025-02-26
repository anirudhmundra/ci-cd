package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}


func indiaHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, India!")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/india", indiaHandler)
	http.ListenAndServe(":8080", nil)
}
