package main

import "fmt"

func modify1(x int) {
	x = 100
}

// x *int： x是整型指针，值是地址
// *x：对指针x指向的地址进行取值操作
func modify2(x *int) {
	*x = 100
}

func ptr() {
	// 指针
	a := 100
	b := &a
	fmt.Println(a, &a)
	fmt.Println(b, &b)

	// 指针取值
	fmt.Println(*b)

	// 指针传值
	c := 200
	modify1(c)
	fmt.Println(c)
	modify2(&c)
	fmt.Println(c)
}
