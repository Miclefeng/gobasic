package main

import (
	"fmt"
)

func maxSubArray(nums []int) int {
	// 分配二维数组保存每次计算结果
	res := make([][]int, len(nums))
	maxNum := nums[0]
	for i := 0; i < len(nums); i++ {
		// 保存开始的元素
		res[i] = append(res[i], nums[i])
		for j := i + 1; j < len(nums); j++ {
			// 保存本次计算结果
			res[i] = append(res[i], res[i][j-i-1]+nums[j])
			if maxNum < res[i][j-i-1]+nums[j] {
				maxNum = res[i][j-i-1] + nums[j]
			}
			fmt.Println(res[i][j-i-1] + nums[j])
		}
		fmt.Println(res[i])
		if maxNum < res[i][0] {
			maxNum = res[i][0]
		}
	}
	return maxNum
}

/*
	这是一个典型的动态规划问题
	用i 做哨兵，根据题目理解，需要找出一个连续的子数组，使得和最大，这时候，哨兵i 要么是前面子数组的一份子，要么是新数组的起始
	而 maxSum 代表了，在循序遍历数组时，左右子数组的和的最大值
	因此 前面子数组的和 我们用 currentSum 表示，就出现了下面的公式
	currentSum = max(a[i], currentSum + a[i])
	maxSum = max(currentSum, maxSum)

	这是典型的动态规划的问题，只要深刻的理解了动态规划，你就发现这就是几行代码搞定的问题
*/
func maxSubArray2(nums []int) int {
	maxSum := nums[0]
	currentSum := 0
	for i := 0; i < len(nums); i++ {
		currentSum = getMax(nums[i]+currentSum, nums[i])
		maxSum = getMax(maxSum, currentSum)
		fmt.Printf("i=%d,num=%2d,cSum=%2d,mSum=%2d\n", i, nums[i], currentSum, maxSum)
	}
	return maxSum
}
func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	//nums = []int{1,4}
	//maxNum := maxSubArray(nums)
	maxNum := maxSubArray2(nums)
	fmt.Println(maxNum)
}
