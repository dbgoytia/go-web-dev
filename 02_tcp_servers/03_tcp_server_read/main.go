package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		// io.WriteString(conn, "\nHello from TCP server")
		// fmt.Fprintln(conn, "\nHow is your day?")
		// fmt.Fprintf(conn, "%v", "Well, I hope!")
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
	}
	defer conn.Close()

	// we never really get here
	// we have an open connection string but incorreclty handled.
	fmt.Println("code got here.")
}
