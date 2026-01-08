package main

import (
	"fmt"
	"sync"
)

var wg1 sync.WaitGroup

func test_channel() {
	var ch1 chan int
	fmt.Println(ch1) //只声明，没初始化，nil

	ch2 := make(chan int, 10)
	fmt.Println(ch2)        //初始化后，是地址类型
	fmt.Printf("%T\n", ch2) //chan int类型

	ch2 <- 100 //将1发送到ch2中
	ch2 <- 200
	close(ch2)

	for v := range ch2 {
		fmt.Println(v)
	}
	fmt.Println("通道已关闭")

	//发送、接受操作
	// 无缓冲通道
	ch3 := make(chan int)
	go recv(ch3)
	ch3 <- 99
	fmt.Println("发送完毕")

	//有缓冲通道
	ch4 := make(chan int, 1)
	wg1.Add(1)
	go recv1(ch4)
	ch4 <- 100
	fmt.Println("发送完毕")
	wg1.Wait()

}

func recv(ch chan int) {
	num := <-ch
	fmt.Println("接收成功", num)
}

func recv1(ch chan int) {
	defer wg1.Done()
	num := <-ch
	fmt.Println("接收成功", num)
}
