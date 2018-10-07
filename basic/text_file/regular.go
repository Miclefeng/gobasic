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

	buf = "3.14 567 ahas 1.23 7. 8.99 lsjjf 6.66"
	reg = regexp.MustCompile(`\d\.[\d]+`)
	result = reg.FindAllStringSubmatch(buf, -1)
	fmt.Println(result)

	html := `
	<div>miclefeng</div>
	<div>golang
		 python
		 php
		 mysql
	</div>
	<div>redis</div>
	<div>kafka</div>
`
	reg = regexp.MustCompile(`<div>(?s:(.*?))</div>`)
	result = reg.FindAllStringSubmatch(html, -1)
	for _, text := range result{
		fmt.Println("Text[1] = ", text[1])
	}
}
