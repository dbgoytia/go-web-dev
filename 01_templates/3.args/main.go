package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	name := os.Args[1]

	tpl := `
		<!DOCTYPE html>
		<html lang="en>
			<head>
				<meta charaset="UTF8">
				<title> Hello World! </title>
			</head>

			<body>
				<h1> ` + name + `</h1>
			</body>
		</html>
	`

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("Error creating file", err)
	}
	defer nf.Close()

	io.Copy(nf, strings.NewReader(tpl))

}
