package main

import "fmt"
import "sync"
import "runtime"

var counter int = 0

func Count(lock *sync.Mutex) {
	lock.Lock()
	counter++
	lock.Unlock()
}

func Add(x, y int) {
	z := x + y
	// 0  4  16  10  12  14  2  6  8  18
	fmt.Print(z, "  ")
}

// 我们在10个goroutine中共享了变量counter。每个goroutine执行完成后，将counter的值加1。因为10个goroutine是并发执行的，所以我们还引入了锁，也就是代码中的lock变量。每次对n的操作，都要先将锁锁住，操作完成后，再将锁打开。在主函数中，使用for循环来不断检查counter的值（同样需要加锁）。当其值达到10时，说明所有goroutine都执行完毕了，这时主函数返回，程序退出。
func main() {
	lock := &sync.Mutex{}
	for i := 0; i < 10; i++ {
		go Add(i, i)
		go Count(lock)
	}

	for {
		lock.Lock()
		c := counter
		lock.Unlock()
		runtime.Gosched()
		if c >= 10 {
			break
		}
	}
}
