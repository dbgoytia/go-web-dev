package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	s := "I felt so good like anything was possible\n I hit cruise control and rubbed my eyes\n The last three days"

	scanner := bufio.NewScanner(strings.NewReader(s))

	// You can implement either built-in split functions,
	// or you can implement your own.
	//
	// scanner.Split(bufio.ScanRunes)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}
