package main

import (
	"io/ioutil"
	"log"
	"strings"
	"strconv"
	"os"
	//"time"
)
/*
Leser tallene fra filen, og legger dem sammen
og skriver resultatet inn en fil igjen.
 */
func SumFromFile() {
	data, err := ioutil.ReadFile("enFil.txt")
	ErrorHandling(err)

	// implementer time.Sleep() for å sjekke at sigint faktisk fungerer
	//time.Sleep(10000000000)

	string1 := string(data)
	stringSlice := strings.SplitAfter(string1, " ")

	a := stringSlice[0]
	b := stringSlice[1]

	//Fjerner mellomrom fra stringen slik at konvertering til int
	//går smertefritt
	c := strings.Replace(a, " ", "", -1)
	d := strings.Replace(b, " ", "", -1)


	//Converter c og d til int
	tall1, err := strconv.Atoi(c)
	ErrorHandling(err)
	tall2, err := strconv.Atoi(d)
	ErrorHandling(err)
	sum := tall1 + tall2

	//Konverterer sum til string slik skrivingen til filen lar seg gjøre
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