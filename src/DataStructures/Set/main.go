package main

import (
	"DataStructures/Set/BinarySearchTreeTSet"
	"DataStructures/Set/FileOperation"
	"fmt"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/4/1 下午4:38
 */

func main()  {
	fmt.Println("Pride and Prejudice")
	words1 := FileOperation.ReadFile("src/DataStructures/Set/pride-and-prejudice.txt")
	fmt.Println("Total words: ", len(words1))
	bstSet := BinarySearchTreeTSet.Instance
	for _, v  := range words1 {
		bstSet.Add(v)
	}
	fmt.Println("Total different words: ", bstSet.GetSize())

	fmt.Println("A Tale of Two Cities")
	words2 := FileOperation.ReadFile("src/DataStructures/Set/a-tale-of-two-cities.txt")
	fmt.Println("Total words: ", len(words2))
	bstSet2 := BinarySearchTreeTSet.Instance
	for _, v  := range words2 {
		bstSet2.Add(v)
	}
	fmt.Println("Total different words: ", bstSet2.GetSize())
}