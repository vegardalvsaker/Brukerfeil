package main

import (
	"os"
	"io/ioutil"
	"fmt"
	"os/signal"
	"syscall"
)
var tall1 = os.Args[1]
var tall2 = os.Args[2]

func main() {
	//Implenterer SIGINT-håndtering
	sigintHandtering2()
	//Kjører første del av oppgave 3b
	createFile()

}
/*createFile oppretter en fil, og skriver inn
to tall fra os.Args
 */

func createFile() {
	file, err := os.Create("enFil.txt")
	ErrorHandling(err)
	defer file.Close()

	//Skriver inn tallene og et mellomrom for å skille tallene
	file.Write([]byte(tall1))
	file.Write([]byte(" "))
	file.Write([]byte(tall2))

	//se filen sumfromfile.go
	SumFromFile()
}
/*
Leser den addisjonen som blir gjort i sumfromfile.go og printer den ut
 */
func ReadFile() {
	data, err := ioutil.ReadFile("enFil.txt")
	ErrorHandling(err)

	string1 := string(data)

	fmt.Println("THE RESULT IS:", string1)
}

func sigintHandtering2() {
	chsig := make(chan os.Signal, 1)
	signal.Notify(chsig, syscall.SIGINT)
	go func() {
		for sig := range chsig {
			switch sig {
			case syscall.SIGINT:
				fmt.Println("\n Sigint received. Shutting down safely")
				os.Exit(0)
			}
		}
	}()
}