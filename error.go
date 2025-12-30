package main

import (
	"errors"
	"fmt"
)

// Info 定义用户信息结构体类型
type Info struct {
	// 可以根据需要添加字段，例如：
	// ID   int64  `json:"id"`
	// Name string `json:"name"`
}

func Query(id int64) (*Info, error) {
	if id == 0 {
		return nil, errors.New("id不能为0")
	}
	return &Info{}, nil
}

func error_test() {
	_, err := Query(0)
	if err != nil {
		fmt.Println(err)
	}
}
