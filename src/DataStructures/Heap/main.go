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

func main()  {
	num := 10
	maxHeap := MaxHeap.Instance
	for i := 0; i < num; i++ {
		rNum := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(100)
		maxHeap.Add(rNum)
	}
	maxHeap.Print()
	fmt.Println("Max:", maxHeap.FindMax())
	fmt.Println("-----------------------------------")
	maxHeap.ExtractMax()
	maxHeap.Print()
}