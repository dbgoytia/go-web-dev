package main

import "fmt"

func main() {
	name := "Diego Goytia"

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

	fmt.Println(tpl)
}
