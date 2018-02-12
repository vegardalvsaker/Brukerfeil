package vegard

import (
	"fmt"
	"bufio"
	"os"
	//"math/rand"
	//"strings"
	"os/signal"
	"syscall"
	"time"
)

func SignalFunc() {

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "kake" {
			fmt.Println("The kake shut the program down")
			os.Exit(0)
		}
		if text == "pølse" {
			fmt.Println("The pølse shut the program down")
			os.Exit(0)
		}
		if text == "hest"{
			fmt.Println("The hest shut the program down")
			os.Exit(0)
		}
	}
}
/*
Denne funksjonen lar brukeren skru av prossesen ved hjelp av CTRL+C og ved å skrive inn "sigint", "sigterm", og "sigquit".
 */
func Signalgreia(text string) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT)

		switch text {
		case "sigint":
			ch <- syscall.SIGINT
		case "sigterm":
			ch <- syscall.SIGTERM
		case "sigquit":
			ch <- syscall.SIGQUIT
		}

	go func() {
		for sig := range ch {

			switch sig {
			case syscall.SIGTERM:
				fmt.Println("sigterm recieved. Shutting down")
				os.Exit(0)
			case syscall.SIGINT:
				fmt.Println("sigint received. Shutting down")
				os.Exit(0)
			case syscall.SIGQUIT:
				fmt.Println("sigquit received. Shutting down")
				os.Exit(0)

			}
		}
	}()
	time.Sleep(time.Minute)
}
func GetInput() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	Signalgreia(text)
}
/*func RandomTall() {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(3)
	fmt.Println(i)
	Signalgreia(i)
}
*/