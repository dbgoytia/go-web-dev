package main

import (
	"html/template"
	"log"
	"os"
)

func main() {

	// loading template
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	// generate new file
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	defer nf.Close()

	// use text/template to write to output with tempalte
	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatalln(err)
	}

}
