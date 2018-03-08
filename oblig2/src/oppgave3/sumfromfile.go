package main

import (
	"io/ioutil"
	"log"
	"strings"
	"strconv"
	"os"
)

func SumFromFile() {
	data, err := ioutil.ReadFile("enFil.txt")
	ErrorHandling(err)

	string1 := string(data)
	array := strings.SplitAfter(string1, " ")

	a := array[0]
	b := array[1]

	c := strings.Replace(a, " ", "", -1)
	d := strings.Replace(b, " ", "", -1)


	//Converter c og d til
	tall1, err := strconv.Atoi(c)
	ErrorHandling(err)
	tall2, err := strconv.Atoi(d)
	ErrorHandling(err)
	sum := tall1 + tall2

	sumstring := strconv.Itoa(sum)

	file, err := os.Create("enFil.txt")
	ErrorHandling(err)
	defer file.Close()

	file.Write([]byte(sumstring))
	ReadFile()
}

func ErrorHandling(err error) {
	if err != nil {
		log.Fatal(err)
	}
}