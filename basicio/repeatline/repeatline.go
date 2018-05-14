package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

//	s := `abc"d"
//kkk
//123134
//kkk
//p
//p`
	counts := make(map[string]int) // key 为 string,value 为 int
	//input := bufio.NewScanner(strings.NewReader(s))
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		//counts[input.Text()]++
		line := input.Text()
		counts[line] = counts[line] + 1 // value不存在，强制转换为 0
		if len(counts) > 3 {
			break
		}
	}

	fmt.Println()
	fmt.Println(counts)
	fmt.Println()
	for line, n := range counts {
		if n > 1 {
			fmt.Println(n, line)
		}
	}
}
