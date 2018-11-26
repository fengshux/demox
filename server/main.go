package main

import (
	"net"
	"fmt"
	"log"
)

func handleConnection( conn net.Conn ) {

	buff := make([]byte, 1024)	
	//conn.Write([]byte("hello world"))

	n , err := conn.Read( buff )
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(string(buff[:n]))

	fmt.Fprintf(conn, "他坏")
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			log.Fatal(err)
			
		}
		go handleConnection(conn)
	}
}
