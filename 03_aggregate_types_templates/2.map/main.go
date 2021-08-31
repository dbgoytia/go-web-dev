package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	// sages := []string{"gandhi", "mlk", "buddha", "jesus", "muhammad"}

	sages := map[string]string{
		"India":    "Ghandi",
		"America":  "MLK",
		"Meditate": "Buddah",
		"Love":     "Jesus",
		"Prophet":  "Muhammad", // Always best practice to add commas
	}

	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}
