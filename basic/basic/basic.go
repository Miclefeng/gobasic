package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	aa = 3
	bb = 4
	ss = "ghi"
)

// 变量的定义
func variableZeroValue() {
	var a int
	var s string
	fmt.Println(a, s)
	fmt.Printf("%d %q\n", a, s)
}

func variableInitiaValue() {
	var a, b int = 3, 4
	var s string = "abc.txt"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

// complex64 的底数是 float32，complex128 的底数是 float64
func euler() {
	fmt.Println(
		cmplx.Exp(1i*math.Pi) + 1)
	fmt.Printf("%.3f\n",
		cmplx.Pow(math.E, 1i*math.Pi)+1)
}

// 类型转换是强制的
func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64((a*a + b*b))))
	fmt.Println(c)
}

func consts() {
	const (
		filename = "abc.txt"
		a, b     = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enums() {
	const (
		cpp = iota // 0开始自增
		_          // 跳过 1
		python
		golang
		php
	)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, python, golang, php)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func bitOpeartion() {
	c := 1 << 10
	fmt.Println(c)
}

func main() {
	x, y := 3, 5
	fmt.Printf("Initiate x, y = %d, %d\n", x, y)
	x, y = y, x
	fmt.Printf("Swap x, y = %d, %d\n", x, y)
	fmt.Println("Hello World!")
	variableZeroValue()
	variableInitiaValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, bb, ss)
	euler()
	triangle()
	consts()
	enums()
	bitOpeartion()
}
