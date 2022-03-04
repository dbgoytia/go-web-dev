package main

import (
	"fmt"
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
		f, h, err := req.FormFile("q")
		HandleError(w, err)
		defer f.Close()

		// local debug
		fmt.Println("\nfile:", f, "\nheader:", h, "\nerr", err)

		// read
		bs, err := ioutil.ReadAll(f)
		HandleError(w, err)

		s = string(bs)

	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := tpl.ExecuteTemplate(w, "index.gohtml", s)
	HandleError(w, err)
	// io.WriteString(w, `
	// <form method="POST" enctype="multipart/form-data">
	// <input type="file" name="q">
	// <input type="submit">
	// </form>
	// <br>`+s)
}

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
