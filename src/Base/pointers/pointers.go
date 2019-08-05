package main

import "fmt"

func main ()  {
	b := 255
	var a *int = &b
	fmt.Printf("Type if a is %T\n", a)
	fmt.Println("address of b is ", a)
	// 符号 * 可以放在一个指针前，如 *intP，那么它将得到这个指针指向地址上所存储的值
	fmt.Println("value of b is ", *a)
	fmt.Println()
	var i1 = 5
	fmt.Printf("An integer: %d, its location in memory: %p\n", i1, &i1)
	var intP *int
	intP = &i1
	fmt.Printf("The value at memory location %p is %d\n", intP, *intP)

	// 不能得到一个文字或常量的地址
	// const i = 5
	// ptr := &i //error: cannot take the address of i
	// ptr2 := &10 //error: cannot take the address of 10
}
