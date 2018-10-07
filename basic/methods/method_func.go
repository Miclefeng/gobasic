package main

import "fmt"

type rectangle struct {
	length int
	width int
}

// 当一个函数有一个值参数时，它只接受一个值参数。
func area (r rectangle)  {
	fmt.Printf("Area func result: %d\n", (r.length * r.width))
}

//当一个方法有一个值接收者时，它可以接受值和指针接收者。
func (r rectangle) area()  {
	fmt.Printf("Area method result: %d\n", (r.length * r.width))
}

// 个接受指针参数的函数只能接受指针
func perimeter (r *rectangle)  {
	fmt.Println("Perimeter func output: ", 2 * (r.length + r.width))
}

// 一个接收者为指针的方法可以接受值接收者和指针接收者
func (r *rectangle) perimeter()  {
	fmt.Println("Perimeter method output: ", 2 * (r.length + r.width))
}

// 非结构体类型的方法
type myInt int

func (a myInt) add (b myInt) myInt {
	return a + b
}

func main ()  {
	r := rectangle{10, 5}
	area(r)

	r.area()
	p := &r // pointer to r
	// p.area() 使用指针接收者 p 调用了接受一个值接收者的方法 area
	// 对于 p.area()，Go 将其解析为 (&p).area()，因为 area 方法必须接受一个值接收者
	p.area()
	fmt.Println()
	l := rectangle{5, 2}
	b := &l // pointer to l
	perimeter(b)
	// 通过一个值接收者 l 调用接受一个指针接收者的 perimeter 方法。
	// 这是合法的，l.perimeter() 这一行将被 Go 解析为 (&l).perimeter()
	b.perimeter()
	l.perimeter()
	fmt.Println()
	num1 := myInt(6)
	num2 := myInt(8)
	sum := num1.add(num2)
	fmt.Println("Sum is ", sum)
}