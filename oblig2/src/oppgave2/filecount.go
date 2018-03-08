package main

import (
	"fmt"
	"bytes"
	"io/ioutil"
	"os"
	"log"
)

func main() {
	fil := `oblig2/src/files/text.txt`

	filbyte, err := ioutil.ReadFile(fil)
	if err != nil {
		log.Fatal(err)
	}

	//Sjekker hvor mange linjer filen har
	lines := 1 //Fordi den første linjen blir ikke telt
	lineSep := []byte{'\n'}
	lines += bytes.Count(filbyte, lineSep)

	//Printer filnavnet og hvor mange linjer filen har
	filnavn, _ := os.Lstat(fil)
	fmt.Println()
	fmt.Println("Information about:", filnavn.Name())
	fmt.Println()
	fmt.Println("Number of lines in file:", lines)
	fmt.Println()

	var runesUsed []byte
	lengde := len(filbyte)

	//Finner aller runes som er i filen
	for i := 0x00; i < 0xFF; i++ {
		for v := 0; v < lengde; v++ {
			if int(filbyte[v]) == i {
				runesUsed = append(runesUsed, filbyte[v])
				v = lengde
			}
		}
	}

	//Oppretter tomt map
	var m= make(map[int]int)
	//Oppretter tom countResult for å legge inn resultatene av telling til runene.
	var countResult = []int{}

	//Teller hvor mange ganger hver rune har blitt brukt og legger den til i et map
	for i := 0; i < len(runesUsed); i++ {
		runeSep := []byte{runesUsed[i]}
		count := bytes.Count(filbyte, runeSep)
		countResult = append(countResult, count)
		m[int(runesUsed[i])] = count
	}
	//Henter ut en sortert liste fra resultatene ovenfor og finner de fem største
	sortedIntSlice := bubbleSort(countResult)
	lengthOfSortedIntSlice := len(sortedIntSlice) - 5
	fiveLargestInts := sortedIntSlice[lengthOfSortedIntSlice:]

	fmt.Println("Most common runes:")
	fmt.Println()
	//Går gjennom mapet for og sjekker alle verdiene opp mot de fem største. Breaker ut av loopen da den finner det den skal finne
Loop1:
	for key, value := range m {
		if value == fiveLargestInts[4] {
			fmt.Printf("1. Rune: %c", key)
			fmt.Printf(", Counts: %d \n", value)
			delete(m, key) //Sletter keyen for å forikre at samme rune ikke kommer flere ganger
			break Loop1 //Forsikrer at printen gjelder kun for en key/value
		}
	}
Loop2:
	for key, value := range m {
		if value == fiveLargestInts[3] {
			fmt.Printf("2. Rune: %c", key)
			fmt.Printf(", Counts: %d \n", value)
			delete(m, key) //Sletter keyen for å forikre at samme rune ikke kommer flere ganger
			break Loop2 //Forsikrer at printen gjelder kun for en key/value
		}
	}
Loop3:
	for key, value := range m {
		if value == fiveLargestInts[2] {
			fmt.Printf("3. Rune: %c", key)
			fmt.Printf(", Counts: %d \n", value)
			delete(m, key)
			break Loop3 //Forsikrer at printen gjelder kun for en key/value
		}
	}
Loop4:
	for key, value := range m {
		if value == fiveLargestInts[1] {
			fmt.Printf("4. Rune: %c", key)
			fmt.Printf(", Counts: %d \n", value)
			delete(m, key)
			break Loop4 //Forsikrer at printen gjelder kun for en key/value
		}
	}
Loop5:
	for key, value := range m {
		if value == fiveLargestInts[0] {
			fmt.Printf("5. Rune: %c", key)
			fmt.Printf(", Counts: %d \n", fiveLargestInts[0])
			delete(m, key)
			break Loop5 //Forsikrer at printen gjelder kun for en key/value
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