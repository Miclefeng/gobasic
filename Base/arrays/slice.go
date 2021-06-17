package main

import (
	"fmt"
	"strings"
)

func removeElement(nums []int, val int) int {
	length := len(nums)
	j := 0
	newNums := []int{}
	for i := 0; i < length; i++ {
		if nums[j] == val {
			newNums = append(newNums, nums[j])
			nums = append(nums[0:j], nums[j+1:]...)
		} else {
			j++
		}
	}
	return len(newNums)
}

func removeElement2(nums []int, val int) int {
	for i := 0; i < len(nums); {
		if nums[i] == val {
			nums = append(nums[0:i], nums[i+1:]...)
		} else {
			i++
		}
	}
	return len(nums)
}

func main() {
	a := make([]int, 0, 1)
	ocap := cap(a)
	for i := 0; i < 20; i++ {
		a = append(a, i)

		if ncap := cap(a); ocap < ncap {
			fmt.Printf("Cap: %d ===> %d\n", ocap, ncap)
			ocap = ncap
		}
	}

	str := "test"
	sl := strings.Split(str, ",")
	fmt.Println(sl[0])
	fmt.Println()
	s2 := []int{0,1,2,2,3,0,4,2}
	s2 = []int{3, 2, 2, 3}
	val := 2
	length := removeElement(s2, val)
	fmt.Println(length)
}
