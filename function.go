package main

import "fmt"

func intSum(x []int) int {
	fmt.Println(x)
	return 0
}

func intSum3(x int, y ...int) int {
	fmt.Println(x, y)
	sum := x
	for _, v := range y {
		sum = sum + v
	}
	return sum
}

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}
func function() {
	// s := []int{1, 2, 3, 4, 5}
	// intSum(s)
	// intSum3(1, 2, 3)

	// 匿名函数
	// sum := func(x int, y int) int {
	// 	return x + y
	// }(1, 2)
	// fmt.Println(sum)

	// sum1 := func(x, y int) int {
	// 	return x + y
	// }
	// fmt.Println(sum1(1, 2))

	// // 闭包
	// add := func(x int) func(y int) int {
	// 	return func(y int) int {
	// 		return x + y
	// 	}
	// }(10)
	// fmt.Println(add(1))

	//defer
	// fmt.Println("start")
	// defer fmt.Println(1)
	// defer fmt.Println(2)
	// defer fmt.Println(3)
	// fmt.Println("end")

	// fmt.Println(f1())
	// fmt.Println(f2())
	// fmt.Println(f3())
	// fmt.Println(f4())
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20
}
