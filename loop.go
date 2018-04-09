package main

import (
	"strconv"
	"fmt"
)

func converToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func main() {
	a := 5
	a /= 2
	fmt.Println(a)
}
