package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

// create a FuncMap to register functions
// here's two example functions, uc, and ft.

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func main() {
	buddah := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs.",
	}

	me := sage{
		Name:  "Diego",
		Motto: "Want a beer?",
	}

	sages := []sage{buddah, me}

	data := struct {
		Wisdom []sage
	}{
		sages,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
