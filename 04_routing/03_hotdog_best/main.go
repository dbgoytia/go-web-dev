package main

import (
	"io"
	"net/http"
)

// Hotdog handler
func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "doggy doggy doggy")
}

// Hotcat handler
func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "kitty kitty kitty")
}

func main() {

	http.HandleFunc("/dog", d) // Attach to default ServeMux
	http.HandleFunc("/cat", c)

	http.ListenAndServe(":8080", nil) // This means we will be using the default handle
}
