package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", upload)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func upload(w http.ResponseWriter, req *http.Request) {

	var s string

	if req.Method == http.MethodPost {
		// open
		file, _, err := req.FormFile("q")
		HandleError(w, err)
		defer file.Close()

		// read
		bs, err := ioutil.ReadAll(file)
		HandleError(w, err)

		s = string(bs)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := tpl.ExecuteTemplate(w, "index.gohtml", s)
	HandleError(w, err)
}

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
