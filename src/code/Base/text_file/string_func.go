package main

import (
	"fmt"
	"strings"
)

func main()  {

	str1 := "seafood"
	// contains 查看字符串是否包含子串，返回bool
	fmt.Println(strings.Contains(str1, "foo"))
	fmt.Println(strings.Contains(str1, "bar"))
	fmt.Println(strings.Contains(str1, ""))
	fmt.Println(strings.Contains("", ""))

	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, ", "))

	// 在字符串中查出子串所在的位置，没有返回-1
	fmt.Println(strings.Index("chicken", "ken"))
	fmt.Println(strings.Index("chicken", "dmr"))

	fmt.Println("ba" + strings.Repeat("na", 2))

	// 在字符串中将 old 字符串替换为 new 字符串， n 表示替换的次数，小于 0 为全部替换
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	fmt.Println(strings.Replace("oink oink oink", "o", "l", -1))

	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Printf("%q\n", strings.Split("a man a plan canal panama", "a"))
	fmt.Printf("%q\n", strings.Split(" xyz ", "")) // [" " "x" "y" "z" " "]

	fmt.Printf("Fields are: %q\n", strings.Fields(" foo bar baz")) // Fields are: ["foo" "bar" "baz"]
}
