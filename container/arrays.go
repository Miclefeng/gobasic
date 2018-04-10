package main

import "fmt"

// 数组是值传递
func printArray(arr *[5]int)  {
	arr[0] = 100
	for i := range arr {
		fmt.Println(arr[i])
	}
}

func main() {
	var arr1 = [5]int{}
	arr2 := [4]int{1, 3, 5, 7}
	arr3 := [...]int{2, 4, 6, 8, 10}
	var grid = [...][2]int{{5, 6}, {1, 2}, {3, 4}}
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}
	for i, v := range arr3 {
		fmt.Println(i, v)
	}
	// _可以在任何地方省略变量
	for _, v := range arr3 {
		fmt.Println(v)
	}
	fmt.Println("printArray(arr1)")
	printArray(&arr1)
	fmt.Println("printArray(arr3)")
	printArray(&arr3)
	fmt.Println("arr1 and arr3")
	fmt.Println(arr1, arr3)
}
