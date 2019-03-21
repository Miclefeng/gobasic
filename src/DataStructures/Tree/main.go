package main

import (
	"DataStructures/Tree/BinarySearchTree"
	"fmt"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/3/8 下午2:44
 */

func main() {
	bst := BinarySearchTree.Instance

	arr := []interface{}{"D", "E", "A", "F", "G", "J", "C"}

	for _, v := range arr {
		bst.NewAdd(v)
	}
	fmt.Println(bst.Search("F"))
	fmt.Println(bst.GetSize())
	fmt.Println("-------------------------")
	bst.PreOrder()
	fmt.Println()
	fmt.Println("-------------------------")
	bst.PreOrderNR()
	fmt.Println()
	fmt.Println("-------------------------")
	bst.InOrder()
	fmt.Println()
	fmt.Println("-------------------------")
	bst.PostOrder()
	fmt.Println()
	fmt.Println("-------------------------")
	bst.LevelOrder()
	fmt.Println()
	fmt.Println("min", bst.Minimum().E)
	fmt.Println("max", bst.Maximum().E)
	bst.RemoveMin()
	fmt.Println("-------------------------")
	bst.InOrder()
	fmt.Println()
	bst.RemoveMax()
	fmt.Println("-------------------------")
	bst.InOrder()
	fmt.Println()
	bst.Remove(5)
	fmt.Println("-------------------------")
	bst.InOrder()
	fmt.Println()
	bst.Remove(4)
	fmt.Println("-------------------------")
	bst.InOrder()
}
