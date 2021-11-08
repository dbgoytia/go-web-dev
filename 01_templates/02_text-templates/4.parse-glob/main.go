package main

import (
	"log"
	"os"
	"text/template"
)

func main() {

	// load all available templates
	tpl, err := template.ParseGlob("templates/*")
	if err != nil {
		log.Fatalln(err)
	}

	// write output to stdout
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	// only execute for one of the available templates
	err = tpl.ExecuteTemplate(os.Stdout, "vespa.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

}
