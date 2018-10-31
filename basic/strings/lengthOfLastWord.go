package main

import "fmt"

func lengthOfLastWord(s string) int {
	start := 0
	ok := true
	for _, v := range s {
		if ' ' == v {
			ok = true
		} else {
			if ok {
				start = 0
			}
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
