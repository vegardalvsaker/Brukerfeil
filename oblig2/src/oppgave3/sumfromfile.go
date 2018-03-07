package main

import (
	"io/ioutil"
	"log"
	"fmt"
	"strings"
	"strconv"
)

func main() {

	data, err := ioutil.ReadFile("enFil.txt")
	if err != nil {
		log.Fatal(err)
	}

	string1 := string(data)
	array := strings.SplitAfter(string1, " ")

	a := array[0]
	b := array[1]

	c := strings.Replace(a, " ", "", -1)
	d := strings.Replace(b, " ", "", -1)



	tall1, err := strconv.Atoi(c)
	if err != nil {
		log.Fatal(err)
	}

	tall2, err := strconv.Atoi(d)
	if err != nil {
		log.Fatal(err)
	}


	sum := tall1 + tall2


	fmt.Println("File sum:", sum)

}
