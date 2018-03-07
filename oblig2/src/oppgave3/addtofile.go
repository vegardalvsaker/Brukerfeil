package main

import (
	"os"
	"log"
)

var tall1 = os.Args[1]
var tall2 = os.Args[2]

func main() {
	file, err := os.Create("enFil.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file.Write([]byte(tall1))
	file.Write([]byte(" "))
	file.Write([]byte(tall2))
	file.Write([]byte(" "))
}
