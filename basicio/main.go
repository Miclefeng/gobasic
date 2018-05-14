package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)

	if n > 0 {
		return p[:n], nil
	}

	return p, err
}

// 从字符串读取
func simpleReadFromString() {
	data, _ := ReadFrom(strings.NewReader("From String"), 12)
	fmt.Println(data)
}

// 从输入流读取
func simpleReadStdin() {
	fmt.Println("Please input from stdin:")
	data, _ := ReadFrom(os.Stdin, 11)
	fmt.Println(data)
}

// 从文件读取
func simpleReadFile() {
	file, _ := os.Open("main.go")
	defer file.Close()

	data, _ := ReadFrom(file, 12)
	fmt.Println(string(data))
}

func main() {
	//simpleReadFromString()
	simpleReadStdin()
	//simpleReadFile()
}
