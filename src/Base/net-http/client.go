package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// 打开链接
	conn, err := net.Dial("tcp", "127.0.0.1:50000")
	if err != nil {
		fmt.Println("Error dialing...", err.Error())
		return
	}

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("First, what`s your name?")
	clientName, _ := inputReader.ReadString('\n')
	trimmedClient := strings.Trim(clientName, "\n")

	for {
		fmt.Println("What to send to the server? Type Q to quit.")
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\n")
		if trimmedInput == "Q" {
			return
		}
		_, err = conn.Write([]byte(trimmedClient + " says: " + trimmedInput))
	}
}
