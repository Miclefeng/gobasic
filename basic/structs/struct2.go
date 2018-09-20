package main

import "fmt"

type mystr string

type person struct {
	name mystr
	sex byte
	age int
}

type student struct {
	person
	int
	mystr
}

func main() {
	s := student{person{"micle", 'm', 26}, 1, "BJ"}
	fmt.Printf("s = %+v\n", s)
	fmt.Println(s.person, s.int, s.mystr)
	s.person = person{"go", 10, 'f'}
	s.int = 2
	s.mystr = "UA"
	fmt.Println(s.name, s.sex, s.age, s.int, s.mystr)
}
