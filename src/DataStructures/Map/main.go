package main

import (
	"DataStructures/Map/BinarySearchTree"
	"DataStructures/Map/LinkedListMap"
	"DataStructures/Set/FileOperation"
	"fmt"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/4/2 上午11:33
 */

func main()  {
	fmt.Println("Pride and Prejudice")

	words2 := FileOperation.ReadFile("src/DataStructures/Set/pride-and-prejudice.txt")
	fmt.Println("Total words: ", len(words2))
	bMap := BinarySearchTreeMap.Instance
	for _, v  := range words2 {
		if bMap.Contains(v) {
			num := bMap.Get(v)
			intN := num.(int)
			bMap.Set(v, intN + 1)
		} else {
			bMap.Add(v, 1)
		}
	}
	fmt.Println("Total different words: ", bMap.GetSize())
	fmt.Println("Frequency of PRIDE: ", bMap.Get("pride"))
	fmt.Println("Frequency of PREJUDICE: ", bMap.Get("prejudice"))

	fmt.Println("----------------------------------")

	words1 := FileOperation.ReadFile("src/DataStructures/Set/pride-and-prejudice.txt")
	fmt.Println("Total words: ", len(words1))
	lMap := LinkedListMap.Instance
	for _, v  := range words1 {
		if lMap.Contains(v) {
			num := lMap.Get(v)
			intN := num.(int)
			lMap.Set(v, intN + 1)
		} else {
			lMap.Add(v, 1)
		}
	}
	fmt.Println("Total different words: ", lMap.GetSize())
	fmt.Println("Frequency of PRIDE: ", lMap.Get("pride"))
	fmt.Println("Frequency of PREJUDICE: ", lMap.Get("prejudice"))
}