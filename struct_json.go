package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	ID     int
	Gender string
	Name   string
}

type class struct {
	Title    string
	Students []Student
}

func struct2json() {
	c := &class{
		Title:    "101",
		Students: make([]Student, 0, 20),
	}

	for i := 0; i < 10; i++ {
		stu := Student{
			ID:     i,
			Gender: "女",
			Name:   fmt.Sprintf("stu%02d", i),
		}
		c.Students = append(c.Students, stu)
	}

	//json序列化
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}
	fmt.Printf("json:%s\n", data)

	//json反序列化
	str := `{"Title":"101","Students":[{"ID":0,"Gender":"女","Name":"stu00"},{"ID":1,"Gender":"女","Name":"stu01"},{"ID":2,"Gender":"女","Name":"stu02"},{"ID":3,"Gender":"女","Name":"stu03"},{"ID":4,"Gender":"女","Name":"stu04"},{"ID":5,"Gender":"女","Name":"stu05"},{"ID":6,"Gender":"女","Name":"stu06"},{"ID":7,"Gender":"女","Name":"stu07"},{"ID":8,"Gender":"女","Name":"stu08"},{"ID":9,"Gender":"女","Name":"stu09"}]}`
	c1 := &class{}
	err = json.Unmarshal([]byte(str), c1)
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return
	}
	fmt.Printf("%#v\n", c1)
}
