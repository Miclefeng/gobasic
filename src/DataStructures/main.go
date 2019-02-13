package main

import (
	"DataStructures/Array"
	"DataStructures/Stack"
	"fmt"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/2/13 下午2:43
 */

func main() {
	array := Array.Instance

	for i :=0;i<10;i++ {
		array.AddLast(i)
	}
	array.Print()
	array.AddFirst(20)
	array.Print()
	array.RemoveFirst()
	array.RemoveFirst()
	array.RemoveFirst()
	array.RemoveFirst()
	array.RemoveFirst()
	array.RemoveElement(7)
	array.Print()

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
