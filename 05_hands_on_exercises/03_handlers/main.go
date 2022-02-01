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

func me(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(w, "name.gohtml", "Diego")
	if err != nil {
		log.Fatalln(err)
	}

}

func cat(w http.ResponseWriter, req *http.Request) {
	cat := animal{
		Name:     "cat",
		Greeting: "bshhh  bshh bshh",
	}
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(w, "animals.gohtml", cat)
	if err != nil {
		log.Fatalln(err)
	}
}

func dog(w http.ResponseWriter, req *http.Request) {
	dog := animal{
		Name:     "dog",
		Greeting: "dog dog doggy",
	}
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(w, "animals.gohtml", dog)
	if err != nil {
		log.Fatalln(err)
	}
}

func root(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(w, "index.gohtml", req.Form)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	http.Handle("/", http.HandlerFunc(root))
	http.Handle("/me/", http.HandlerFunc(me))
	http.Handle("/dog/", http.HandlerFunc(dog))
	http.Handle("/cat/", http.HandlerFunc(cat))
	http.ListenAndServe(":8080", nil)
}
