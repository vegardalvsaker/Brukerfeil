package main
import (
"fmt"
"net"
"bufio"
	//"os"
	"os"
)

func main() {
	conn, err := net.Dial("udp", "localhost:17")
	if err != nil {
		fmt.Printf("Some error %v", err)
		return
	}

	defer conn.Close()
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Press enter to recieve quote of the day")
		text, _  :=reader.ReadString('\n')
		fmt.Fprintf(conn, text + "\n")
		message, _ := bufio.NewReader(conn).ReadString('!')
		fmt.Println("Server:\n" + message)


}