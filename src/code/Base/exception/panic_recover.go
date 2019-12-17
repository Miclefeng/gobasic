package main

import "fmt"

func testa() {
	fmt.Println("AAAAAAAAAA")
}

func testb(x int)  {
	// defer 函数在函数结束后执行，多个defer(使用栈进行存储)遵循先进后出原则
	defer func() {
		// recover() 可以打印panic中的错误信息，恢复panic退出的程序
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