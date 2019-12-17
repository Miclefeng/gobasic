package main

import (
	"fmt"
)

func tryRecover()  {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error occured : ", err)
		} else {
			panic(r)
		}
	}()

	//panic(errors.New("This is an error!"))
	//b := 0
	//a := 5 / b
	//fmt.Println(a)
	panic(
		fmt.Sprintf("I don't know what to do %v", 123))
}

func main() {
	tryRecover()
}
