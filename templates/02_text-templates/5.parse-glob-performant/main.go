package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

// init runs only once when the program is first executed
func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	// write output to stdout
	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	// only execute for one of the available templates
	err = tpl.ExecuteTemplate(os.Stdout, "vespa.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

}
