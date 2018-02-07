package main

import "fmt"
import "os"
import "os/signal"
import "syscall"
import "time"


func main() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		for sig := range ch {
			switch sig {
			case syscall.SIGTERM:
				fmt.Println("sigterm recieved")
				os.Exit(0)
			case syscall.SIGINT:
				fmt.Println("sigint received. Shutting down")
				os.Exit(0)
			case syscall.SIGQUIT:
				fmt.Println("sigquit received")
				os.Exit(0)



			}
		}
	}()
	time.Sleep(time.Minute)
}
