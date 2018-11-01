package main

import "fmt"

func plusOne(digits []int) []int {
	mark := false
	end := len(digits) - 1
	for i := end; i >=0; i-- {
		if i == end && 9 == digits[i] {
			mark = true
		}
		if mark == true && digits[i] + 1 >= 10 {
			digits[i] = 0
			if 0 == digits[i] && 0 == i {
				num := []int{1}
				digits = append(num, digits...)
			}
		} else {
			mark = false
			digits[i]++
			break
		}
	}

	return digits
}

func plusOne2(digits []int) []int {
	add := 0
	end := len(digits) - 1
	digits[end] += 1
	for i := end; i >= 0; i-- {
		digits[i] += add
		if digits[i] >= 10 {
			digits[i] = digits[i] % 10
			add = 1
		} else {
			add = 0
		}
		if 0 == i && 1 == add {
			digits = append([]int{1}, digits...)
		}
	}
	return digits
}

func main() {
	//digits := []int{4, 3, 2, 1}
	digits := []int{9, 9, 9, 9}
	num := plusOne2(digits)
	fmt.Println(num)
}
