package main

import (
	"fmt"
	"sync"
)

/*互斥锁*/
var x int
var wg2 sync.WaitGroup
var m sync.Mutex

func add() {
	for i := 0; i < 5000; i++ {
		m.Lock()
		x = x + 1
		m.Unlock()
	}
	wg2.Done()
}
func test_mutex() {
	wg2.Add(2)
	go add()
	go add()
	wg2.Wait()
	fmt.Println(x)
}
