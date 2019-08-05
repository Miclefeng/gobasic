package main

import (
	"fmt"
	"math"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	return mergeCommonPrefix(strs, 0, len(strs) - 1)
}

func mergeCommonPrefix(strs []string, l, r int) string {
	if l >= r {
		return strs[l]
	}

	mid := (l + r) / 2
	left := mergeCommonPrefix(strs, l, mid)
	right := mergeCommonPrefix(strs, mid+1, r)
	return mergeCommon(left, right)
}

func mergeCommon(left , right string) string {
	min := math.Min(float64(len(left)), float64(len(right)))
	for i := 0; i < int(min); i++ {
		if left[i] != right[i] {
			return left[0:i]
		}
	}
	return left[0:int(min)]
}

func main()  {
	arr := []string{"flower","flow","flight"}
	str := longestCommonPrefix(arr)
	fmt.Println(str)
}
