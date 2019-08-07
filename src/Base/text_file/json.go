package main

import (
	"encoding/json"
	"fmt"
)

type IT struct {
	Company string
	Subjects []string
	IsOk bool
	Price float64
}

func main ()  {

	// 通过结构体生成json文本
	s := IT{"cooler", []string{"Go", "PHP", "Python"}, true, 666.666}

	// 将切片转换为json文本
	res, err := json.Marshal(s)
	// 格式化json文本
	res, err = json.MarshalIndent(s, "", " ")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("res = ", string(res))

	// 通过map生成json文本
	m := make(map[string]interface{}, 4)
	m["company"] = "cooler"
	m["subjects"] = []string{"Go", "PHP", "Python"}
	m["isok"] = true
	m["price"] = 666.666

	res, err = json.MarshalIndent(m, "", " ")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("res = ", string(res))
}
