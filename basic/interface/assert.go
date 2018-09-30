package main

import "fmt"

type Student struct {
	name string
	id int
}

func main ()  {
	// 空接口可以保持任何类型的变量
	i := make([]interface{}, 3)
	i[0] = 1
	i[1] = "Hello World"
	i[2] = Student{"Miclefeng", 26}

	for i, data := range i {
		//
		if value, ok := data.(int); ok == true {
			fmt.Printf("I[%d] 类型为int，内容为%d\n", i, value)
		} else if value, ok := data.(string); ok == true {
			fmt.Printf("I[%d] 类型为string，内容为%s\n", i, value)
		} else if value, ok := data.(Student); ok == true {
			fmt.Printf("I[%d] 类型为Student，内容为%v\n", i, value)
			fmt.Printf("I[%d] 类型为Student，内容为name = %s, id = %d\n", i, value.name, value.id)
		}
		switch value := data.(type) {
		case int:
			fmt.Printf("I[%d] 类型为int，内容为%d\n", i, value)
		case string:
			fmt.Printf("I[%d] 类型为string，内容为%s\n", i, value)
		case Student:
			fmt.Printf("I[%d] 类型为Student，内容为%v\n", i, value)
			fmt.Printf("I[%d] 类型为Student，内容为name = %s, id = %d\n", i, value.name, value.id)
		}
		fmt.Println()
	}
}
