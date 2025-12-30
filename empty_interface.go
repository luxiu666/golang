package main

import "fmt"

/*
空接口
1 空接口的定义
     空接口是指没有定义任何方法的接口类型。
     因此任何类型都可以视为实现了空接口。也正是因为空接口类型的这个特性，空接口类型的变量可以存储任意类型的值。
2 空接口的应用
     2.1 空接口作为函数的参数
	     使用空接口实现可以接收任意类型的函数参数
	 2.2 空接口作为map的值
	     使用空接口实现可以保存任意值的字典
*/

type Empty interface{}
type Cow struct{}

// 使用空接口实现可以接收任意类型的函数参数。
func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}

func empty_interface() {
	var e Empty

	e = "hello"
	fmt.Printf("type:%T value:%v\n", e, e)

	e = 100
	fmt.Printf("type:%T value:%v\n", e, e)

	e = true
	fmt.Printf("type:%T value:%v\n", e, e)

	e = Cow{}
	fmt.Printf("type:%T value:%v\n", e, e)

	show(1000)
	show("你好")
	show(false)

	var studentInof = make(map[string]interface{})
	studentInof["name"] = "张三"
	studentInof["age"] = 18
	studentInof["isMale"] = true
	fmt.Println(studentInof)

}
