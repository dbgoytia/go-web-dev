package main

import (
	"log"
	"os"
	"text/template"
)

type sage struct {
	Name  string
	Motto string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	buddah := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs.",
	}

	err := tpl.Execute(os.Stdout, buddah)
	if err != nil {
		log.Fatalln(err)
	}
}
