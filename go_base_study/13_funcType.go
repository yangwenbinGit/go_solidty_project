package main

import "fmt"

// go语言种函数也是一种类型 使用funcType进行定义 注意的是在定义类型的时候func后面没有函数名
type funcType func(int, int) int

func Calculate(a, b int, f funcType) int {
	return f(a, b)
}

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func Mul(a, b int) int {
	return a * b
}

func main() {
	a := Calculate(10, 5, Add)
	b := Calculate(10, 5, Sub)
	c := Calculate(10, 5, Mul)
	// 匿名函数
	d := Calculate(10, 5, func(a, b int) int {
		return a / b
	})
	// 闭包
	e := Calculate(10, 5, func(a, b int) int {
		return a % b
	})
	fmt.Println(a, b, c, d, e)
}
