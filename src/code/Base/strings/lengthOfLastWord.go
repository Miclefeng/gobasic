package main

import "fmt"

func lengthOfLastWord(s string) int {
	start := 0
	// ' ' 标记
	ok := true
	for _, v := range s {
		if ' ' == v {
			// 标记为 true
			ok = true
		} else {
			// 如果' '之后有字符，长度赋值为0
			if ok {
				start = 0
			}
			// 标记为 false
			ok = false
			start++
		}
	}
	return start
}

func main() {
	s := "hello world str"
	length := lengthOfLastWord(s)
	fmt.Println(length)
}
