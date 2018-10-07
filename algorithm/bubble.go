package main

import "fmt"

func main() {
	arr := [...]int{11, 13, 6, 9, 4, 8, 16}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr) - i - 1; j++ {
			if arr[j] > arr[j + 1] {
				arr[j], arr[j + 1] = arr[j + 1], arr[j]
			}
		}
		fmt.Println(arr)
	}
	fmt.Println(arr)
}
