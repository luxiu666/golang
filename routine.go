package main

import (
	"fmt"
	"sync"
)

/*启动1个goroutine*/
var wg sync.WaitGroup

func hello() {
	fmt.Println("hello world")
	wg.Done() //告知当前goroutine完成，-1
}
func routine_test() {
	wg.Add(1) //添加一个goroutine，+1
	go hello()
	fmt.Println("main goroutine")
	wg.Wait() //等待所有goroutine完成
	fmt.Println("main goroutine done")
}

/*启动多个goroutine*/
func hellos(i int) {
	defer wg.Done() //函数返回前执行，推荐写法
	fmt.Println("hello world", i)
}

func test_routines() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go hellos(i)
	}
	wg.Wait()
}
