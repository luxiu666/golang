package main

import "fmt"

// global variables
var (
	name string = "hayri"
	age  int    = 21
)

func variables() {
	// local variables
	n := 10
	m := 100
	fmt.Println(name, age)
	fmt.Println(n, m)
}
