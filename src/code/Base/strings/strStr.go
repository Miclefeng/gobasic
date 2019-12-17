package main

import "fmt"

func strStr(haystack string, needle string) int {
	index := -1
	nLen := len(needle)
	hLen := len(haystack)
	if 0 == nLen {
		return 0
	}

	if 0 == hLen {
		return -1
	}

	for i := 0; i <= hLen - nLen; i++ {
		fmt.Println(haystack[i:nLen+i])
		// 一直偏移 needle 长度的切片
		if needle == haystack[i:nLen+i] {
			index = i
			break
		}
	}
	return index
}

func main() {
	haystack := "aaa"
	needle := "a"
	index := strStr(haystack, needle)
	fmt.Println(index)
}
