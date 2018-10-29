package main

import (
	"fmt"
	"net"
)

func doServerStuff(conn net.Conn)  {
	for {
		buf := make([]byte, 512)
		length, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading...", err.Error())
			return
		}
		fmt.Printf("Received data: %v\n", string(buf[:length]))
	}
}

func main()  {
	fmt.Println("starting the server...")
	listener, err := net.Listen("tcp", "0.0.0.0:50000")
	if err != nil {
		fmt.Println("Error listening...", err.Error())
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting...", err.Error())
			return
		}

		go doServerStuff(conn)
	}
}
