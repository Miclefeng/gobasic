package main

import (
	"fmt"
	"strings"
	"io"
	"bufio"
)

type intGen func() int  // 函数类型

// 1, 1, 2, 3, 5, 8 ...
// a, b
//    a, b
//       a, b
func fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a + b
		return a
	}
}

func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 100 {
		return 0, io.EOF // 文件读到末尾
	}
	// TODO: incorrect if p is too small!
	// 把 intGen，strings.NewReader 通过 struct 存储起来
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader)  {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	f := fibonacci()
	//fmt.Println(f()) // 1 a=1, b=1
	//fmt.Println(f()) // 1 a=1, b=2
	//fmt.Println(f()) // 2 a=2, b=3
	//fmt.Println(f()) // 3 a=3, b=5
	//fmt.Println(f()) // 5 a=5, b=8
	//fmt.Println(f()) // 8 a=8, b=13
	//fmt.Println(f()) // 13 a=13, b=21
	//fmt.Println(f()) // 21 a=21, b=34
 	//fmt.Println(f()) // 34 a=34, b=55
 	printFileContents(f)
}
