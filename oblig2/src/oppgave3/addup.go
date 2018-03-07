package main

import (
	"fmt"
	"time"
)

var ch = make(chan int, 2)

var tall1 int = 4
var tall2 int = 8

func main() {
	//Funksjon A
	go func() {
		ch <- tall1
		ch <- tall2
		time.Sleep(1000)
		fmt.Println(<- ch)
		fmt.Println("Ferdig")
	}()

	//Funksjon B
	go func() {
		add1 := <- ch
		add2 := <- ch
		sum := add1 + add2
		ch <- sum
	}()

	time.Sleep(10000000)
}
