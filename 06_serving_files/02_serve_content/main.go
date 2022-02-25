package main

import (
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/dog.png", dog)
	http.HandleFunc("/dog-again.png", dogTwo)
	http.ListenAndServe(":8080", nil) // We use nil to simply attach to the default servemux
}

func dog(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("dog.png")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}

	http.ServeContent(w, req, fi.Name(), fi.ModTime(), f)
}

func dogTwo(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "dog.png")
}
