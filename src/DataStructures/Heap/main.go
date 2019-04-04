package main

import (
	"DataStructures/Heap/MaxHeap"
	"fmt"
	"math/rand"
	"time"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/4/3 上午10:24
 */

func main() {
	//num := 10
	//maxHeap := MaxHeap.Instance
	//for i := 0; i < num; i++ {
	//	rNum := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(100)
	//	maxHeap.Add(rNum)
	//}
	//maxHeap.Print()
	//fmt.Println("Max:", maxHeap.FindMax())
	//fmt.Println("-----------------------------------")
	//maxHeap.ExtractMax()
	//maxHeap.Print()
	//maxHeap.Replace(66)
	//maxHeap.Print()
	//fmt.Println("-----------------------------------")
	//arr := []interface{}{23, 44, 66, 33, 55, 77, 99, 88}
	//maxHeap.Heapify(arr)
	//maxHeap.Print()

	newNum := 20000
	arr := make([]interface{}, newNum)
	for i := 0; i < newNum; i++ {
		rNum := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(10000000)
		arr[i] = rNum
	}

	time2 := testHeap(arr, false)
	fmt.Println("Without heapify:", time2, "s")

	time1 := testHeap(arr, true)
	fmt.Printf("With heapify: %f s", time1)
}

func testHeap(data []interface{}, heapify bool) float64 {
	sTime := time.Now().UnixNano()

	mHeap := MaxHeap.Instance
	if heapify {
		mHeap.Heapify(data)
	} else {
		for _, v := range data {
			mHeap.Add(v)
		}
	}

	arr := make([]int, len(data))
	for i := 0; i < len(data); i++ {
		arr[i] = mHeap.ExtractMax().(int)
	}

	for i := 1; i < len(data); i++ {
		if arr[i-1] < arr[i] {
			panic("Error")
		}
	}
	fmt.Println("Test MaxHeap completed.")

	eTime := time.Now().UnixNano()

	return float64(eTime-sTime) / 1000000000;
}
