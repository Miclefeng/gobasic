package main

import "fmt"

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


}
