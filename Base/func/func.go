package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("Unsupported operation: %s\n", op)

	}
}

//func div(a, b int) (int, int) {
//	return a / b, a % b
//}

func div(a, b int) (q, r int) {
	return a / b, a % b
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling func %s with args "+"(%d, %d)\n", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

// 可变参数列表
func sumArgs(num ...int) int {
	sum := 0
	for i := range num {
		sum += num[i]
	}
	return sum
}

// go语言是值传递
//func swap(a, b *int)  {
//	*b, *a = *a, *b
//}
func swap(a, b int) (int, int) {
	return b, a
}

/*
var cache Cache  <==>  func f(cache Cache)
引用传递
*/

func main() {
	if result, err := eval(3, 4, "x"); err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(result)
	}
	q, r := div(13, 4)
	fmt.Println(q, r)
	fmt.Println(apply(pow, 3, 4))
	fmt.Println(apply(
		func(i int, i2 int) int {
			return int(math.Pow(float64(i), float64(i2)))
		}, 3, 4))
	fmt.Println(sumArgs(1, 2, 3, 4, 5))
	a, b := 3, 4
	a, b = swap(a, b)
	fmt.Println(a, b)
}
