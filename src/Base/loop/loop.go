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
		fmt.Printf("%d\t", lsb)
		// int 转换成 int字符串
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func main() {
	fmt.Println(
		//converToBin(5),
		converToBin(13),
		//converToBin(1024),
		converToBin(0),
	)

	//printFile("abc.txt")
	s := `abc"d"
kkk
123134
p`
	printFileContents(strings.NewReader(s))

	// i 一直为 0，所以会无线循环下去
	//for i := 0; i < 3; {
	//	fmt.Println("Value of i:", i)
	//}

	for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaa"; i, j,
		s = i+1, j+1, s + "a" {
		fmt.Println("Value of i, j, s:", i, j, s)
	}
	fmt.Println()
	// 循环到 s1 = "aaaaa" 为止
	s1 := ""
	for ; s1 != "aaaaa"; {
		fmt.Println("Value of s:", s1)
		s1 = s1 + "a"
	}


LABEL1:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				// continue 语句指向 LABEL1，当执行到该语句的时候，就会跳转到 LABEL1 标签的位置。
				continue LABEL1
				// break，则不会只退出内层循环，而是直接退出外层循环了
				// break LABEL1
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}
}
