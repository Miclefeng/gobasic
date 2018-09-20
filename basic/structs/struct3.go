package main

import "fmt"

type Person struct {
	name string
	gender byte
	age int
}

type Student struct {
	*Person
	id int
	addr string
}

func main() {
	s1 := Student{&Person{"micle", 'm', 26}, 1, "BJ"}
	fmt.Println(s1.name, s1.gender, s1.age, s1.id, s1.addr)

	var s2 Student // 先定义变量
	s2.Person = new(Person) // 分配内存空间
	s2.name = "go"
	s2.gender = 'f'
	s2.age = 11
	s2.id = 2
	s2.addr = "UA"
	fmt.Println(s2.name, s2.gender, s2.age, s2.id, s2.addr)
}
