package main

import (
	"os"
	"fmt"
	"bufio"
)

func main() {
	fmt.Println(os.Args)
	fmt.Println()
	if len(os.Args) < 2 {
		fmt.Println("Args less than 2")
		return
	}

	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err, "err1")
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	var line int
	for {
		_, isPrefix, err := reader.ReadLine()

		if err != nil {
			fmt.Println(err, "err2")
			break
		}

		if !isPrefix {
			line++
		}
	}
	fmt.Println(line)
}
