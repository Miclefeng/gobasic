package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var res []int
	for n := 0; n < len(nums); n++ {
		for j := n + 1; j < len(nums); j++ {
			if nums[n] + nums[j] == target {
				res = append(res, n, j)
			}
		}
	}
	return res
}

func twoSum2(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	res := make([]int, 2)
	for i, v := range nums {
		w := target - v
		if j, ok := m[w]; ok {
			fmt.Println(j, ok)
			res[0] = j
			res[1] = i
			return res
		} else {
			m[v] = i
		}
	}
	return res
}

func main()  {
	var nums  = []int{2, 7, 11, 15}
	target := 9
	res := twoSum2(nums, target)
	fmt.Println(res)
}