package fib

// 1, 1, 2, 3, 5, 8 ...
// a, b
//    a, b
//       a, b
func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a + b
		return a
	}
}