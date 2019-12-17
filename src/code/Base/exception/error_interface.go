package main

import (
	"errors"
	"fmt"
)

func MyDiv(a, b int) (result int, err error)  {
	if 0 == b {
		err = errors.New("除数不能为0")
		return
	}
	result = a / b
	return
}

func main()  {
	err1 := fmt.Errorf("%s", "This is normal err1!")
	fmt.Println(err1)

	err2 := errors.New("This is normal err2!")
	fmt.Println(err2)

	res, err := MyDiv(20, 10)
	if err != nil {
		fmt.Println("Err = ", err)
	} else {
		fmt.Println("Res = ", res)
	}
}
