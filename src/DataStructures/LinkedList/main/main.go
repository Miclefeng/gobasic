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
	fmt.Println(linkedList.GetFirst())
	for i := 0; i < 10; i++ {
		fmt.Println(linkedList.Get(i))
	}
	fmt.Println(linkedList.GetLast())
}
