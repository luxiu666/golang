package main

import "fmt"

func slice() {
	/*切片声明*/
	var a []string
	var b = []int{}
	var c = []bool{false, true}
	//var d = []bool{false, true}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(a == nil)
	fmt.Println(b == nil)
	fmt.Println(c == nil)

	/*切片构造*/
	var a1 = [5]int{1, 2, 3, 4, 5}
	s := a1[1:3] // s := a[low:high]
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))
	s2 := s[3:4] // 索引的上限是cap(s)而不是len(s)
	fmt.Printf("s2:%v len(s2):%v cap(s2):%v\n", s2, len(s2), cap(s2))

	//make
	var a2 = make([]int, 2, 10)
	fmt.Printf("a2:%v len(a2):%v cap(a2):%v\n", a2, len(a2), cap(a2))

	//append
	var s3 []int
	s3 = append(s3, 1)
	s3 = append(s3, 2, 3, 4)
	s4 := []int{5, 6, 7}
	s3 = append(s3, s4...)
	fmt.Println(s3)

	//copy slice
	cs := []int{1, 2, 3, 4, 5}
	cs2 := make([]int, 5, 5)
	copy(cs2, cs)
	fmt.Println(cs2)
	fmt.Println(cs)
	cs[0] = 100
	fmt.Println(cs2)
	fmt.Println(cs)

}
