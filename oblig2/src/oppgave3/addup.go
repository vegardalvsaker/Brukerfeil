package main

import (
	"fmt"
	"time"
	"os"
	"os/signal"
	"syscall"
	"bufio"
	"strconv"
	"strings"
	"log"
)

var ch = make(chan int, 2)

func main() {
	//Aktiverer sigInt-håndtering
	sigintHandtering()

	//Leser inn to tall fra terminal
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a number: ")
	tall1Str, err1 := reader.ReadString('\n')
	errHandling(err1)
	fmt.Print("Enter one more number: ")
	tall2Str, err2 := reader.ReadString('\n')
	errHandling(err2)

	//Forsikrer at programmet fungerer også på Windows
	c := strings.Replace(tall1Str, "\r", "", -1)
	d := strings.Replace(tall2Str, "\r", "", -1)

	e := strings.Replace(c, "\n", "", -1)
	f := strings.Replace(d, "\n", "", -1)

	//Gjør inputen fra stdin om til int slik at vi kan addere
	tall1, err3 := strconv.Atoi(e)
	errHandling(err3)

	tall2, err4 := strconv.Atoi(f)
	errHandling(err4)

	oppgave3a(tall1, tall2)
}

func oppgave3a(tall1, tall2 int) {
	//Funksjon A
	go func() {
		ch <- tall1
		ch <- tall2
		time.Sleep(1)
		fmt.Println(<- ch)
		fmt.Println("Ferdig")
	}()

	//Funksjon B
	go func() {
		add1 := <- ch
		add2 := <- ch
		sum := add1 + add2
		ch <- sum
	}()
	time.Sleep(10000000)
}

//Håndterer errors
func errHandling(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//Håndterer en eventuell sigint
func sigintHandtering() {
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