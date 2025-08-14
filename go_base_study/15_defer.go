package main

import "fmt"

//TODO defer用于延迟一个函数的调用 如果多个defer 采用的是先进后出的原则  打印结果为
// main函数开始
// defer2
// defer1
// main函数结束

func main() {

	fmt.Println("main函数开始")
	defer fmt.Println("main函数结束")
	defer fmt.Println("defer1")
	defer fmt.Println("defer2")

	// defer常和匿名函数结合使用
	a, b := 10, 20
	defer func(x int) {
		fmt.Println(x, b)
	}(a + b)
	fmt.Println(a, b)

}
