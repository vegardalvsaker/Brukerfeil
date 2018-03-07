package main

import (
	"fmt"
	"bytes"
	"io/ioutil"
	"os"
	"log"
)

func main() {
	//fil := os.Args[1]

	fil := `C:\Users\Vegard\Documents\dev\Go\is-105\Brukerfeil\oblig2\src\oppgave2\ronny.txt`
	filbyte, err := ioutil.ReadFile(fil)
	if err != nil {
		log.Fatal(err)
	}

	//Sjekker hvor mange linjer filen har
	lines := 0
	lineSep := []byte{'\n'}
	lines += bytes.Count(filbyte, lineSep)

	//Printer filnavnet og hvor mange linjer filen har
	filnavn, _ := os.Lstat(fil)
	fmt.Println("Information about:", filnavn.Name())
	fmt.Println("Number of lines in file:", lines)

	var runesUsed []byte
	lengde := len(filbyte)

	//Finner aller runes som er i filen
	for i := 0x00; i < 0x80; i++ {
		for v := 0; v < lengde; v++ {
			if int(filbyte[v]) == i {
				runesUsed = append(runesUsed, filbyte[v])
				v = lengde
			}
		}
	}

	//Oppretter tomt map
	var m= make(map[int]int)
	//Oppretter tom intSlice for å legge inn resultatene av telling til runene.
	var intSlice = []int{}

	//Teller hvor mange ganger hver rune har blitt brukt og legger den til i et map
	for i := 0; i < len(runesUsed); i++ {
		runeSep := []byte{runesUsed[i]}
		count := bytes.Count(filbyte, runeSep)
		intSlice = append(intSlice, count)
		m[int(runesUsed[i])] = count
	}
	//Henter ut en sortert liste fra resultatene ovenfor og finner de fem største
	sortedIntSlice := bubbleSort(intSlice)
	lengthOfSortedIntSlice := len(sortedIntSlice) - 5
	fiveLargestInts := sortedIntSlice[lengthOfSortedIntSlice:]

	//Går gjennom mapet for og sjekker alle verdiene opp mot de fem største. Breaker ut av loopen da den finner det den skal finne
Loop:
	for key, value := range m {
		if value == fiveLargestInts[4] {
			fmt.Printf("1. Rune: %q", key)
			fmt.Printf(", Counts: %d \n", fiveLargestInts[4])
			break Loop
		}
	}
Loop2:
	for key, value := range m {
		if value == fiveLargestInts[3] {
			fmt.Printf("2. Rune: %q", key)
			fmt.Printf(", Counts: %d \n", fiveLargestInts[3])
			break Loop2
		}
	}
Loop3:
	for key, value := range m {
		if value == fiveLargestInts[2] {
			fmt.Printf("3. Rune: %q", key)
			fmt.Printf(", Counts: %d \n", fiveLargestInts[2])
			break Loop3
		}
	}
Loop4:
	for key, value := range m {
		if value == fiveLargestInts[1] {
			fmt.Printf("4. Rune: %q", key)
			fmt.Printf(", Counts: %d \n", fiveLargestInts[1])
			break Loop4
		}
	}
Loop5:
	for key, value := range m {
		if value == fiveLargestInts[0] {
			fmt.Printf("5. Rune: %q", key)
			fmt.Printf(", Counts: %d \n", fiveLargestInts[0])
			break Loop5
		}
	}
}

func bubbleSort(list []int) []int {
	// find the length of list n
	n := len(list)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if list[j] > list[j+1] {
				temp := list[j+1]
				list[j+1] = list[j]
				list[j] = temp
			}
		}
	}
	return list
}
