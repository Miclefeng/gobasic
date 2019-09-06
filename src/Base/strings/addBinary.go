package main

import (
	"fmt"
	"strconv"
)

func addBinary(a string, b string) string {
	length := 0
	if len(a) > len(b) {
		length = len(a)
	} else {
		length = len(b)
	}
	//next := '0'
	c := ""
	for i := length -1; i >= 0; i-- {
		if a[i] == b[i] && '1' == a[i] {
			c += "0"
			//next = '1'
		} else {

		}
	}
	return ""
}

func main() {
	//var a uint32
	//a = 1
	//b := "1"
	//c := addBinary(a, b)
	str := "3C2B8A1A952102E0";
	int64str, err := strconv.ParseUint(str, 16, 64)
	fmt.Println(int64str, err)
	//fmt.Println(a >> 1)
}
