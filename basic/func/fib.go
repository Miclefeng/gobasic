package main

import (
	"fmt"
	"time"
)

const LIM  = 41

var fibs[LIM] int64

func fibonacci(n int) (res int64) {
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}

	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n - 1) + fibonacci(n - 2)
	}
	fibs[n] = res
	return 
}

func main()  {
	var result int64
	start := time.Now()
	for i := 0; i < LIM; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is %d\n", i, result)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Println(start.UnixNano())
	fmt.Println(end.UnixNano())
	fmt.Printf("Calculation is %d\n", delta)
}