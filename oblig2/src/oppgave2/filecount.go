package main

import (
	"fmt"
	"bytes"
	"io/ioutil"
	"os"
)
func main() {
	//fil := os.Args[1]
	fil := `C:\Users\Vegard\Documents\dev\Go\is-105\Brukerfeil\oblig2\src\oppgave2\ronny.txt`
	filnavn,_ := os.Lstat(fil)
	fmt.Println("Information about:", filnavn.Name())
	filbyte, _ := ioutil.ReadFile(fil)
	lengde := len(filbyte)
	//r, err := os.Open(fil)
	fmt.Print(filbyte, lengde)
	fmt.Println(int(filbyte[0]))
	fmt.Println(int(filbyte[1]))
	fmt.Println(int(filbyte[2]))
	fmt.Println(int(filbyte[3]))
	fmt.Println(int(filbyte[4]))

	lines := 0
	lineSep := []byte{'\n'}

	lines += bytes.Count(filbyte, lineSep)

    //test := hex.Dump(filbyte)

	fmt.Println("Number of lines in file:", lines)


	var runes []byte
	var foundRunes []byte

	for i := 0x00; i < 0x80; i++ {

		for v := 0; v < lengde; v++ {
			if int(filbyte[v]) == i {
				foundRunes = append(runes, filbyte[v])
				v = lengde
			}
		}
	//fmt.Printf("%q", foundRunes)
	}
	fmt.Println(foundRunes)

	for i := 0; i < len(foundRunes); i++ {

		runeSep := []byte {foundRunes[i]}
		bytes.Count(filbyte, runeSep)

	}













	/*fmt.Println("Most common runes:")
	fmt.Println("1. Rune:",,"Counts:",)
	fmt.Println("2. Rune:",,"Counts:",)
	fmt.Println("3. Rune:",,"Counts:",)
	fmt.Println("4. Rune:",,"Counts:",)
	fmt.Println("5. Rune:",,"Counts:",)
*/}