package main

import (
	"os"
	"log"
	"io/ioutil"
	"fmt"
)

var tall1 = os.Args[1]
var tall2 = os.Args[2]


func createFile() {
	file, err := os.Create("enFil.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file.Write([]byte(tall1))
	file.Write([]byte(" "))
	file.Write([]byte(tall2))

	SumFromFile()
}

func ReadFile() {
	data, err := ioutil.ReadFile("enFil.txt")
	ErrorHandling(err)

	string1 := string(data)

	fmt.Println("THE RESULT IS:", string1)
}

func main() {
	createFile()
}
