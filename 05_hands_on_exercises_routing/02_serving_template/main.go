package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type animal struct {
	Name     string
	Greeting string
}

func me(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(res, "name.gohtml", "Diego")
	if err != nil {
		log.Fatalln(err)
	}

}

func cat(res http.ResponseWriter, req *http.Request) {
	cat := animal{
		Name:     "cat",
		Greeting: "bshhh  bshh bshh",
	}
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(res, "animals.gohtml", cat)
	if err != nil {
		log.Fatalln(err)
	}
}

func dog(res http.ResponseWriter, req *http.Request) {
	dog := animal{
		Name:     "dog",
		Greeting: "dog dog doggy",
	}
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(res, "animals.gohtml", dog)
	if err != nil {
		log.Fatalln(err)
	}
}

func root(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(res, "index.gohtml", req.Form)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/me/", me)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/cat/", cat)
	http.ListenAndServe(":8080", nil)
}
