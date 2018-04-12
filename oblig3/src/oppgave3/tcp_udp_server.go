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
	go TCPserver()
	go UDPserver()
	for    {
		time.Sleep(100000000)
	}
}

func TCPserver() {
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
	ServerConn, err3 := net.ListenUDP("udp", ServerAddr)
	checkError(err3)
	defer ServerConn.Close()
	buf := make([]byte, 1024)
		for {
			_, addr, _ := ServerConn.ReadFromUDP(buf)
			ServerConn.WriteToUDP([]byte("Quote of the day: \n" + quote), addr)
		}
}