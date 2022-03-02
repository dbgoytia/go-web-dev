package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/toby", dog)
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err) // Handle erros on the server
}

func dog(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("toby.png")
	if err != nil {
		http.Error(w, "file not found", http.StatusNotFound) // http.Error handling
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		http.Error(w, "file not found", http.StatusNotFound)
		return
	}
	http.ServeContent(w, req, f.Name(), fi.ModTime(), f)
}
