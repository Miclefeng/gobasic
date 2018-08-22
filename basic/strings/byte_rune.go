package main

import (
	"fmt"
	"unicode/utf8"
)

func printBytes(s string)  {
	for i := 0;i < len(s);i++ {
		fmt.Printf("%x\t", s[i])
	}
}

func printChars(s string)  {
	// rune 是 Go 中的内置类型，它是 int32 的别名。
	// 在 Go 中，rune 表示一个 Unicode 码点。
	// 无论一个码点会被编码为多少个字节，它都可以表示为一个 rune
	runes := []rune(s)
	for i := 0;i < len(runes);i++ {
		fmt.Printf("%c\t", runes[i])
	}
}

func printCharsAndBytes(s string) {
	// range for 遍历字符串
	for index, chr := range s {
		fmt.Printf("%c starts at byte %d\n", chr, index)
	}
}

func utf8length (s string)  {
	// utf8 包 提供了 func RuneCountInString(s string) (n int) 来获取字符串的长度
	// RuneCountInString 返回字符串中 Unicode 字符的个数
	// len 返回字符串中 byte 的个数
	fmt.Printf("length of %s is %d\n", s, utf8.RuneCountInString(s))
}

func mutate (s []rune) string {
	s[0] = 'M'
	return string(s)
}

func main() {
	byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	str := string(byteSlice)
	fmt.Println(str)
	fmt.Println()
	name := "Señor"
	printBytes(name)
	fmt.Println()
	printChars(name)
	fmt.Println()
	printCharsAndBytes(name)
	fmt.Println()
	utf8length(name)
	name2 := "Pets"
	utf8length(name2)
	h := "hello"
	// 改变一个字符串中的字符，我们需要先把字符串转换为 rune 切片
	// 然后修改切片中的内容，最后将这个切片转换回字符串
	fmt.Println(mutate([]rune(h)))
	fmt.Println(h)
}
