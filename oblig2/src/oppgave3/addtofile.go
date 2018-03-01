package main

import (
	"fmt"
	"log"
)

func main() {
	ch := make(chan int, 1)
	//scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter a number")
	input, err := fmt.Scan()
	if err != nil {
		log.Fatal(err)
		}
	ch <- input

	fmt.Println(input)


}
