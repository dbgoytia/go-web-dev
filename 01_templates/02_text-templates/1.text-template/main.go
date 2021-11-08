package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("tpl.gohtml") // variadic function, you can parse N html's.
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
