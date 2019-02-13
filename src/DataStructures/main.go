package main

import (
	"DataStructures/Stack"
	"fmt"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/2/13 下午2:43
 */

func main() {
	stack := Stack.Instance

	for i := 0;i < 10; i++ {
		stack.Push(i)
	}
	stack.Print()
	stack.Pop()
	stack.Print()
	stack.Pop()
	stack.Pop()
	stack.Pop()
	stack.Print()
	fmt.Println(stack.Top())
}
