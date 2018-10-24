package main

import (
	"fmt"
	"time"
)

func sendData(ch chan string) {
	ch <- "start"
	ch <- "golang"
	ch <- "php"
	ch <- "python"
	ch <- "end"
}

func getData(ch chan string) {
	var input string
	for {
		input = <-ch
		fmt.Printf("%s ", input)
	}
}

func main()  {
	ch := make(chan string)
	go sendData(ch)
	go getData(ch)
	time.Sleep(1e9)
}