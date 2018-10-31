package main

import "fmt"

func searchInsert(nums []int, target int) int {
	for i, v := range nums {
		if target <= v {
			return i
		}
		if i == len(nums) - 1 && v < target {
			return i + 1
		}
	}
	return 0
}

func main()  {
	nums := []int{1, 3, 5, 6}
	target := 7
	index := searchInsert(nums, target)
	fmt.Println(index)
}
