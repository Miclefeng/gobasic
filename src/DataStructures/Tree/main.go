package main

import (
	"DataStructures/Tree/BinarySearchTree"
	"fmt"
	"strconv"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/3/8 下午2:44
 */

func main() {
	bst := BinarySearchTree.Instance

	arr := []int{5, 3, 8, 2, 6, 4, 9}

	for _, v := range arr {
		bst.NewAdd(v, strconv.Itoa(v))
	}
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

}
