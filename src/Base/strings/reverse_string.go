package main

import "fmt"

func Reverse (s string) string {
	r := []rune(s)
	for i, j := 0, len(r) - 1; i < len(r) / 2; i, j = i + 1, j - 1 {
		// r[i], r[j] 同时赋值，不会原值不会被覆盖
		r[i], r[j] = r[j], r[i]
		// r[i] = r[j] 分开赋值，原值会被覆盖
		// r[j] = r[i]
	}
	return string(r)
}

func main ()  {
	fmt.Println(Reverse("!oG ,olleH"))
}
