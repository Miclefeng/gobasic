package main

import (
	"code/DataStructures/Old/Map/BinarySearchTree"
	"code/DataStructures/Old/Map/LinkedListMap"
	"code/DataStructures/Old/Map/Map"
	"code/DataStructures/Old/Set/FileOperation"
	"fmt"
	"time"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/4/2 上午11:33
 */

func main()  {
	fmt.Println("Pride and Prejudice")

	//words2 := FileOperation.ReadFile("src/DataStructures/Set/pride-and-prejudice.txt")
	//fmt.Println("Total words: ", len(words2))
	//bMap := BinarySearchTreeMap.Instance
	//for _, v  := range words2 {
	//	if bMap.Contains(v) {
	//		num := bMap.Get(v)
	//		intN := num.(int)
	//		bMap.Set(v, intN + 1)
	//	} else {
	//		bMap.Add(v, 1)
	//	}
	//}
	//fmt.Println("Total different words: ", bMap.GetSize())
	//fmt.Println("Frequency of PRIDE: ", bMap.Get("pride"))
	//fmt.Println("Frequency of PREJUDICE: ", bMap.Get("prejudice"))
	//
	//fmt.Println("----------------------------------")
	//
	//words1 := FileOperation.ReadFile("src/DataStructures/Set/pride-and-prejudice.txt")
	//fmt.Println("Total words: ", len(words1))
	//lMap := LinkedListMap.Instance
	//for _, v  := range words1 {
	//	if lMap.Contains(v) {
	//		num := lMap.Get(v)
	//		intN := num.(int)
	//		lMap.Set(v, intN + 1)
	//	} else {
	//		lMap.Add(v, 1)
	//	}
	//}
	//fmt.Println("Total different words: ", lMap.GetSize())
	//fmt.Println("Frequency of PRIDE: ", lMap.Get("pride"))
	//fmt.Println("Frequency of PREJUDICE: ", lMap.Get("prejudice"))

	filename := "src/code/DataStructures/Set/pride-and-prejudice.txt"
	var uMap Map.Map

	bMap := BinarySearchTreeMap.Instance
	uMap = bMap
	bTime := testMap(uMap, filename)
	fmt.Println("BinarySearchTree Map:", bTime, "s")

	fmt.Println("----------------------------------")

	lMap := LinkedListMap.Instance
	uMap = lMap
	lTime := testMap(uMap, filename)
	fmt.Println("LinkedLists Map:", lTime, "s")
}

func testMap(uMap Map.Map, filename string) float64 {
	sTime := time.Now().UnixNano()
	fmt.Println("FileName: ", filename)
	words := FileOperation.ReadFile(filename)
	fmt.Println("Total words: ", len(words))
	for _, v  := range words {
		if uMap.Contains(v) {
			num := uMap.Get(v)
			intN := num.(int)
			uMap.Set(v, intN + 1)
		} else {
			uMap.Add(v, 1)
		}
	}
	fmt.Println("Total different words: ", uMap.GetSize())
	fmt.Println("Frequency of PRIDE: ", uMap.Get("pride"))
	fmt.Println("Frequency of PREJUDICE: ", uMap.Get("prejudice"))
	eTime := time.Now().UnixNano()

	return (float64(eTime) - float64(sTime)) / 1000000000
}