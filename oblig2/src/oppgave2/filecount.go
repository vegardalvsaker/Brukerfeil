package main

import (
	"fmt"
	"bytes"
	"io/ioutil"
)
func main() {
	//fil := os.Args[1]
	fil := `C:\Users\Vegard\Documents\dev\Go\is-105\Brukerfeil\oblig2\src\oppgave2\ronny.txt`
	fmt.Println("Information about:", fil)
	filbyte, _ := ioutil.ReadFile(fil)
	lengde := len(filbyte)
	//r, err := os.Open(fil)

	lines := 0
	lineSep := []byte{'\n'}

	lines += bytes.Count(filbyte, lineSep)

    //test := hex.Dump(filbyte)

	fmt.Println("Number of lines in file:", lines)
	char1x := 0




	for i := 0x00; i < 0x80; i++ {
		for v := 0; v < lengde; v++ {
			if int(filbyte[v]) == i {
				sep := []byte{filbyte[v]}
				char1x += bytes.Count(filbyte, sep)
				//char1 := filbyte[v]
				fmt.Println(char1x)
			}


		}
	}


	//fmt.Println("Most common runes:")


/*
	fmt.Println("1. Rune:",,"Counts:",)
	fmt.Println("2. Rune:",,"Counts:",)
	fmt.Println("3. Rune:",,"Counts:",)
	fmt.Println("4. Rune:",,"Counts:",)
	fmt.Println("5. Rune:",,"Counts:",)
*/}