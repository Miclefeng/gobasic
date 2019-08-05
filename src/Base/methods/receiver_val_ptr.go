package main

import "fmt"

type Employee struct {
	name string
	age  int
}

type address struct {
	city, state string
}

type person struct {
	firstName, lastName string
	address
}

func (a address) fullAddress() {
	fmt.Printf("FullAddress: %s, %s\n", a.city, a.state)
}

// 值接收者
func (e Employee) changeName(newName string) {
	e.name = newName
}

// 指针接收者
// 指针接收者可用于对接收者的修改应该对调用者可以见的场合
// 指针接收者也可用于拷贝结构体代价较大的场合
// 使用指针接收者将避免结构体的拷贝，而仅仅是指向结构体指针的拷贝
func (e *Employee) chageAge(newAge int) {
	e.age = newAge
}

func main() {
	e := Employee{"micelfeng", 26}
	fmt.Printf("Employee name before change: %s\n", e.name)
	e.changeName("Micle")
	fmt.Printf("Employee name after change: %s\n", e.name)
	fmt.Println()
	fmt.Printf("Employee age before change: %d\n", e.age)
	// 因为 changeAge 有一个指针类型的接收者我们必须使用 (&e) 来调用
	// Go允许我们省略 & 符号，因此可以只写为 e.changeAge(51)
	// Go 将 e.changeAge(51) 解析为 (&e).changeAge(51)
	(&e).chageAge(27)
	fmt.Printf("Employee age after change: %d\n", e.age)
	fmt.Println()

	// 匿名字段的方法可以被包含该匿名字段的结构体的变量调用
	// 就好像该匿名字段的方法属于包含该字段的结构体一样。
	p := person{
		firstName: "Micle",
		lastName:  "zss",
		address: address{
			city:  "BeiJ",
			state: "wangjing",
		},
	}
	p.fullAddress()
}
