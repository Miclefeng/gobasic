package main

import "fmt"

func testa() {
	fmt.Println("AAAAAAAAAA")
}

func testb(x int)  {
	// defer 函数在函数结束后执行，多个defer(使用栈进行存储)遵循先进后出原则
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	var a [10]int
	a[x] = 111
	fmt.Println("BBBBBBBBBB")
}

func testc()  {
	fmt.Println("CCCCCCCCCC")
}

func main()  {
	testa()
	testb(9)
	testc()
}