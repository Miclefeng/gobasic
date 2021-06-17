package main

import (
	"fmt"
	"os"
	"bufio"
	"miclefeng/learngo/functional/fib"
)

func tryDefer()  {
	defer fmt.Println(1) // defer 相当于栈，先进后出
	defer fmt.Println(2)
	fmt.Println(3)
	return // return 之后也可以执行 defer
	fmt.Println(4)
}

func tryDefer2()  {
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if 30 == i {
			panic("Printed too many !")
		}
	}
}

func writeFile(filename string)  {
	// os.O_EXCL 检查文件是否存在
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%s, %s, %s\n", pathError.Op, pathError.Path, pathError.Err)
		}
		return
	}

	defer file.Close()

	writer := bufio.NewWriter(file) // 现将文件写的内存中，在写入磁盘文件中
	defer writer.Flush() // 将缓冲中的字节流写入文件

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	//tryDefer2()
	//tryDefer()
	writeFile("fib.txt")
}
