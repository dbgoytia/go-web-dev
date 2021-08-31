package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

// init runs only once when the program is first executed
func init() {
	tpl = template.Must(template.ParseGlob("tpl.gohtml"))
}

func main() {
	// only execute for one of the available templates
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", `Release self-focus; embrace other-focus.`)
	if err != nil {
		log.Fatalln(err)
	}

}
