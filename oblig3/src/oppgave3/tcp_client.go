package main

import (
	"net"
	"log"
	"bufio"
	"fmt"
)

func main() {

	// connect to this socket
	conn, err := net.Dial("tcp", "localhost:17")
	if err != nil {
		log.Fatal(err)
		}
	defer conn.Close()
	// listen for reply
	message, _ := bufio.NewReader(conn).ReadString('!')
	fmt.Print("Server:\n"+message)
}