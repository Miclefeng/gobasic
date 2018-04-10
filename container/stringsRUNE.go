package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes我爱慕课网!" // utf-8
	fmt.Println(len(s)) // len 获取字节数
	fmt.Printf("%s\n", []byte(s))
	fmt.Printf("%x\n", []byte(s))
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()
	for i, ch := range s { // ch is a rune,对utf-8解码,在进行unicode
		fmt.Printf("(%d %X) ", i, ch)
	}
	fmt.Println()

	// 返回字符的长度
	fmt.Println("Rune count: ", utf8.RuneCountInString(s))

	bytes := []byte(s) // []byte 获取所有的字节
	ch, size := utf8.DecodeRune(bytes)
	fmt.Println(size, ch, bytes)
	fmt.Println()
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes) // 只返回第一个字符和字符所占长度
		bytes = bytes[size:] // 分片往后截取字符,每次去除第一个字符
		//fmt.Println(size, ch, bytes)
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	for i, ch := range []rune(s) { // rune 每个字符占4个字节
		fmt.Printf("(%d %c) ", i, ch)
	}
	fmt.Println()
}
