package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// Connect to the server
	conn, err := net.Dial("tcp", "localhost:8080")

	if err != nil {
		println(err)
		return

	}

	// Send some data to the server
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter your message :")
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)

		if userInput == "Exit" {
			println("Exiting .......")
			break
		}
		_, err = conn.Write([]byte(userInput))
		if err != nil {
			println(err)
			return
		}

	}

	defer conn.Close()
}
