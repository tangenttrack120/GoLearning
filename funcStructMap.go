package main

import "fmt"

func calc(x int, y int, op string) (result int) {
	if op == "add" {
		result = x + y
	} else if op == "sub" {
		result = x - y
	} else if op == "mul" {
		result = x * y
	} else if op == "div" {
		result = x / y
	} else {
		fmt.Println("Invalid operation")
		result = 0
	}
	return result
}

func multreslt(x int, y string) (result int, greet string) {
	result = x * 2
	greet = "Hello " + y
	return
}

func recursion(x int) (result int) {
	if x == 0 {
		result = 1
	} else {
		result = x * recursion(x - 1)
	}
	return result
}

// struct
type Person struct {
	name string
	age int
	job string
	salary int32
}

func main() {
	/*
	result := calc(10, 5, "add")
	mult, greet := multreslt(2, "Richard")
	_, groot := multreslt(2, "groot")
	x, _ := multreslt(2, "you")*/
	recurse := recursion(5)
	fmt.Println(recurse)
}