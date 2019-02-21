package main

import (
	"DataStructures/LinkedList"
	"fmt"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/2/21 上午10:20
 */

func main() {
	linkedList := LinkedList.Instance

	for i := 0; i < 10; i++ {
		linkedList.AddLast(i)
	}

	fmt.Println(linkedList.GetSize())
	fmt.Println()
	for i := 0; i < linkedList.GetSize(); i++ {
		fmt.Println(linkedList.Get(i))
	}

	linkedList.Set(5, "a")
	fmt.Println(linkedList.Get(5))
	fmt.Println(linkedList.Contains("a"))
	fmt.Println(linkedList.Contains("b"))
	fmt.Println("---------------------")
	fmt.Println(linkedList.RemoveFirst())
	fmt.Println(linkedList.RemoveLast())
	linkedList.RemoveElement("a")
	fmt.Println("---------------------")
	for i := 0; i < linkedList.GetSize(); i++ {
		fmt.Println(linkedList.Get(i))
	}
}
