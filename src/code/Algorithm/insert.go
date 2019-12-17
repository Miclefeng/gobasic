package main

import "fmt"

func main() {
	arr := [...]int{11, 13, 6, 9, 4, 8, 16}

	for i := 1; i < len(arr); i++ {
		for j := i; j > 0 && arr[j] < arr[j - 1]; j-- {
			arr[j], arr[j - 1] = arr[j - 1], arr[j]
		}
		fmt.Println(arr)
	}
	fmt.Println(arr)
}
