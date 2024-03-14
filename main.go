package main

import (
	"fmt"
	"log"
	"net"
	"os"

	connection "github.com/franklin-01/simple-http-server/connection"
)

func main() {

	flags := connection.ParseFlags(os.Args)
	directory, ok := flags["directory"]

	if ok {
		log.Printf("Serving from: %s", directory)
	} else {
		log.Printf("Serving from: %s", ".")
	}

	listener, error := net.Listen("tcp", "0.0.0.0:4221")
	if error != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	for {

		conn, error := listener.Accept()
		if error != nil {
			log.Fatal(error)
		}
		go connection.HandleConnection(conn, directory)
	}

}
