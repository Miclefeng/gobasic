package main

import "fmt"

func main() {
	arr := [...]int{11, 13, 6, 9, 4, 8, 16}

	for i := 0; i < len(arr); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[min], arr[i] = arr[i], arr[min]
		fmt.Println(arr)
	}
	fmt.Println(arr)
}
