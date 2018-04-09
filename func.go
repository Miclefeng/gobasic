package main

import "fmt"

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

func main() {
	if result, err := eval(3, 4, "x"); err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(result)
	}
	q, r := div(13, 4)
	fmt.Println(q, r)
}
