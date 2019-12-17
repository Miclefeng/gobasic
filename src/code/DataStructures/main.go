package main

import (
	"DataStructures/Array"
	"DataStructures/Queue/ArrayQueue"
	"DataStructures/Queue/LoopQueue"
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

	arrayQueue := ArrayQueue.Instance

	for i :=0; i< 20;i++ {
		arrayQueue.EnQueue(i)
	}
	arrayQueue.Print()

	for i :=0;i < 15;i++ {
		arrayQueue.DeQueue()
	}
	arrayQueue.Print()
	fmt.Println(arrayQueue.GetFront())

	fmt.Println()
	loopQueue := LoopQueue.Instance
	for i :=0; i< 8; i++ {
		loopQueue.EnQueue(i)
	}
	loopQueue.Print()
	loopQueue.DeQueue()
	loopQueue.DeQueue()
	loopQueue.DeQueue()
	loopQueue.Print()
	loopQueue.EnQueue(8)
	loopQueue.EnQueue(9)
	loopQueue.EnQueue(10)
	loopQueue.EnQueue(11)
	loopQueue.Print()
	loopQueue.EnQueue(12)
	loopQueue.EnQueue(13)
	loopQueue.Print()

	arr := []interface{}{8, 5, 3, 6, 9, 7, 15}
	arr[3], arr[4] = arr[4], arr[3]
	fmt.Println(arr)
}
