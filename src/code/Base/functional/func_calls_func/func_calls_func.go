package main

// global scope
var a string

func f1()  {
	// local scope f1()
	a := "O"
	print(a)
	f2()
}

// 变量是在函数定义时确定作用域，不是函数调用时
func f2()  {
	// use global scope
	print(a)
}

func main()  {
	a = "G"
	print(a)
	f1()
}

