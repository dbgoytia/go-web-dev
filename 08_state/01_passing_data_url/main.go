package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q") // search for the value "q" in the query string.
	io.WriteString(w, fmt.Sprintf("My search: %v", v))
}

// curl 'http://localhost:8080/?q=dog'
// My search: dog%
