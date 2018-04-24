package main

import (
	"net"
	"fmt"
	"time"
)

const quote  = "Det skal godt gjøres\nå bare spise en!"

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Launching server...")
	//Kjører serverene i go-routines for å kunne kjøre de samtidig.
	go TCPserver()
	go UDPserver()
	for    {
		time.Sleep(100000000)
	}
}

func TCPserver() {
	//Lytter for tcp-tilkobling på port 17 på localhost
	ln, err := net.Listen("tcp", ":17")
	checkError(err)
	conn, err1 := ln.Accept()
	checkError(err1)
	defer conn.Close()
	for {
		conn.Write([]byte("Quote of the day: \n" + quote))
	}
}

func UDPserver() {
	ServerAddr,err2 := net.ResolveUDPAddr("udp",":17")
	checkError(err2)
	//Lytter for udp-tilkobling på port 17 på localhost
	ServerConn, err3 := net.ListenUDP("udp", ServerAddr)
	checkError(err3)
	defer ServerConn.Close()
	buf := make([]byte, 1024)
		for {
			_, addr, err := ServerConn.ReadFromUDP(buf)
			checkError(err)
			ServerConn.WriteToUDP([]byte("Quote of the day: \n" + quote), addr)
		}
}