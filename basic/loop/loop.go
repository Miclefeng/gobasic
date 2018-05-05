package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"io"
	"strings"
)

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	printFileContents(file)
}

func printFileContents(reader io.Reader)  {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() { // 相当于while]
		fmt.Println(scanner.Text())
	}
}

func converToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func main() {
	fmt.Println(
		converToBin(5),
		converToBin(13),
		converToBin(1024),
		converToBin(0),
	)

	printFile("abc.txt")
	s := `abc"d"
kkk
123134
p`
	printFileContents(strings.NewReader(s))
}