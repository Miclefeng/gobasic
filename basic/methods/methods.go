package main

import (
	"math"
	"fmt"
)

type Rectangle struct {
	length int
	width int
}

type Circle struct {
	radius float64
}

// go 不支持 class 类型。因此通过在一个类型上建立方法来实现与 class 相似的行为
// 同名方法可以定义在不同的类型上，但是 Go 不允许同名函数
func (r Rectangle) Area() int {
	return r.length * r.width
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func main ()  {
	r := Rectangle{10, 5}
	fmt.Printf("Area of rectangle %d\n", r.Area())
	c := Circle{12}
	fmt.Printf("Area of Circle %f\n", c.Area())
}
