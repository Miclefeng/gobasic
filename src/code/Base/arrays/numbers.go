package main

import (
	"time"
	"math/rand"
	"fmt"
)

func createNum(number *int)  {
	seed := time.Now().UnixNano()
	rand.Seed(seed)
	var num int
	for {
		num = rand.Intn(10000)
		if num >= 1000 && num <= 9999 {
			break
		}
	}
	*number = num
}

func getNum(s []int, num int) {
	s[0] = num / 1000
	s[1] = num % 1000 / 100
	s[2] = num % 100 / 10
	s[3] = num % 10
}

func onGame(sysSlice []int)  {
	var num int
	for {
		for {
			fmt.Printf("请输入一个4位的随机数：")
			fmt.Scan(&num)
			if num > 999 && num < 10000 {
				break
			}
			fmt.Println("输入的数不符合要求！")
		}

		intputSlice := make([]int, 4)
		getNum(intputSlice, num)

		correct := 0
		for n := 0; n < len(intputSlice); n++ {
			if intputSlice[n] > sysSlice[n] {
				fmt.Printf("输入的第%d位数大了一些\n", n + 1)
			} else if intputSlice[n] < sysSlice[n] {
				fmt.Printf("输入的第%d位数小了一些\n", n + 1)
			} else {
				fmt.Printf("输入的第%d位数正确\n", n + 1)
				correct++
			}
		}
		if 4 == correct {
			fmt.Println("全部猜对了！")
			break
		}
	}

}

func main() {
	var randNum int
	createNum(&randNum)
	//fmt.Println(randNum)
	sysSlice := make([]int, 4)
	getNum(sysSlice, randNum)
	onGame(sysSlice)
}
