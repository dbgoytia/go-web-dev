package main

import (
	"io"
	"net/http"
)

func me(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Diego")
}

func dog(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog doggy")
}

func root(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "landing zone")
}

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/me/", me)
	http.HandleFunc("/dog/", dog)
	http.ListenAndServe(":8080", nil)
}
