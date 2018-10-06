package main

import (
	"fmt"
	"regexp"
)

func main ()  {
	buf := "abc azc a7c aac 888 a9c tac"
	// 解析规则，解析正则表达式，成功返回解释器
	reg := regexp.MustCompile(`a.c`)

	// 根据规则提取关键信息
	result := reg.FindAllStringSubmatch(buf, -1)
	fmt.Println(result)
}
