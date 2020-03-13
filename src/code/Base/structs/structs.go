package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// Employee 是一个具名结构体（named structure）
// 因为它创建了一个具有名字的结构体类型： Employee
type Employee struct {
	firstName, lastName string
	age, salary int
}

type TagType struct {
	answer bool   `tag:"An important answer"`
	name string `tag:"The name of the thing"`
	price int    `tag:"How much there are"`
}

type IT struct {
	Company  string
	Subjects []string
	IsOk     bool
	Price    float64
}

type University struct {
	Name string `json:"-"` //大学名称
	IsFullTime bool     `json:"isFullTime,string"` //是否全日制
	Subjects   []string `json:"subjects"`          //主要学科
	Tuition    float64  `json:"price, omitempty"`  //学费
}

func main ()  {
	emp1 := Employee{
		firstName: "Micle",
		age: 26,
		salary: 800,
		lastName: "zss",
	}

	emp2 := Employee{"Pangzi", "San", 27, 900}

	fmt.Println("Employee 1 is ", emp1)
	fmt.Println("Employee 2 is ", emp2)
	fmt.Println()

	// 匿名结构体（anonymous structures）
	emp3 := struct {
		firstName, lastName string
		age, salary int
	}{"wei", "zhang", 24, 700}
	fmt.Println("Employee 3 is ", emp3)

	// 结构体变量的 0 值
	var emp4 Employee
	fmt.Println("Employee 4 is ", emp4)

	emp5 := Employee{
		firstName:"Gouzi",
		lastName:"Lianger",
	}
	fmt.Println("Employee 5 is ", emp5)

	emp6 := Employee{"Sam", "Anderson", 55, 6000}
	fmt.Println("First Name:", emp6.firstName)
	fmt.Println("Last Name:", emp6.lastName)
	fmt.Println("Age:", emp6.age)
	fmt.Printf("Salary: $%d\n", emp6.salary)
	fmt.Println()

	tt := TagType{true, "micle", 26}
	for i := 0; i < 3; i++ {
		ttType := reflect.TypeOf(tt)
		ixField := ttType.Field(i)
		fmt.Printf("%v\n", ixField.Tag.Get("tag"))
	}

	fmt.Println()
	t1 := IT{"tencent", []string{"develop", "unittest", "production", "operative"}, false, 12000.0}
	b, err := json.Marshal(t1)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))
	//案例2： json.MarshalIndent的使用,用缩进对输出进行格式化,效果更加json化
	t2 := IT{"HuaWei", []string{"develop", "unittest", "production", "operative"}, false, 13000.0}
	b2, err := json.MarshalIndent(t2, "", "   ")
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b2))
}
