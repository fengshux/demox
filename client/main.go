package main

import (
	"net"
	"fmt"
	"log"
	//"bufio"
)


func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		// handle error
		log.Fatal(err)
	}
	fmt.Fprintf(conn, "olhle")
	//	status, err := bufio.NewReader(conn).ReadString('\n')

	handleConnection(conn)
	

	
}


func handleConnection( conn net.Conn ) {

	buff := make([]byte, 1024)	
	//conn.Write([]byte("hello world"))

	n , err := conn.Read( buff )
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(buff[:n]))	
}
