package main

import (
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

}
