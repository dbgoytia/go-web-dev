package main

import (
	"html/template"
	"log"
	"net/http"
)

// Load the available template (Globs)
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

// Serve the template
type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(w, "index.gohtml", req.Form)
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
