package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "This is an useage of a string"
	fmt.Printf("T/F? Does the string '%s' have prefix %s? ", str, "Th")
	fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))

	fmt.Printf("T/F? Does the string '%s' have suffix %s? ", str, "ing")
	fmt.Printf("%t\n", strings.HasSuffix(str, "ing"))

	fmt.Printf("T/F? Does the string '%s' have contains %s? ", str, "useage")
	fmt.Printf("%t\n", strings.Contains(str, "useage"))

	// Index 返回字符串 str 在字符串 s 中的索引（str 的第一个字符的索引），-1 表示字符串 s 不包含字符串 str：
	// strings.Index(s, str string) int

	// LastIndex 返回字符串 str 在字符串 s 中最后出现位置的索引（str 的第一个字符的索引），-1 表示字符串 s 不包含字符串 str：
	// strings.LastIndex(s, str string) int

	// 如果 ch 是非 ASCII 编码的字符，建议使用以下函数来对字符进行定位：
	// strings.IndexRune(s string, r rune) int

	// eplace 用于将字符串 str 中的前 n 个字符串 old 替换为字符串 new，并返回一个新的字符串，如果 n = -1 则替换所有字符串 old 为字符串 new：
	// strings.Replace(str, old, new, n) string

	// Count 用于计算字符串 str 在字符串 s 中出现的非重叠次数：
	// strings.Count(s, str string) int

	// Repeat 用于重复 count 次字符串 s 并返回一个新的字符串：
	// strings.Repeat(s, count int) string

	// ToLower 将字符串中的 Unicode 字符全部转换为相应的小写字符：
	// strings.ToLower(s) string
	// ToUpper 将字符串中的 Unicode 字符全部转换为相应的大写字符：
	// strings.ToUpper(s) string

	// strings.TrimSpace(s) 来剔除字符串开头和结尾的空白符号；
	// 如果你想要剔除指定字符，则可以使用 strings.Trim(s, "cut") 来将开头和结尾的 cut 去除掉
	// 使用 TrimLeft 或者 TrimRight 剔除开头或者结尾的字符串

	// strings.Fields(s) 将会利用 1 个或多个空白符号来作为动态长度的分隔符将字符串分割成若干小块，并返回一个 slice，
	// 如果字符串只包含空白符号，则返回一个长度为 0 的 slice。

	// strings.Split(s, sep) 用于自定义分割符号来对指定字符串进行分割，同样返回 slice。

	// Join 用于将元素类型为 string 的 slice 使用分割符号来拼接组成一个字符串：
	// strings.Join(sl []string, sep string) string
	fmt.Println()
	str2 := "The quick brown fox jumps over the lazy dog"
	sl := strings.Fields(str2)
	fmt.Printf("Splitted in slice: %v\n", sl)
	for _, val := range sl {
		fmt.Printf("%s - ", val)
	}
	fmt.Println()
	str3 := "GO1|The ABC of Go|25"
	sl2 := strings.Split(str3, "|")
	fmt.Printf("Splitted in slice: %v\n", sl2)
	for _, val := range sl2 {
		fmt.Printf("%s - ", val)
	}
	fmt.Println()
	str4 := strings.Join(sl2, ";")
	fmt.Printf("sl2 joined by ';' : %s\n", str4)
}
