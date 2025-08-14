package main

import "fmt"

// 函数
func test() (int, string) {
	return 1, "hello world"
}

// int默认的值为0
func main() {

	var v1 int = 20

	var v2 = 30

	v3 := 40

	var v4 int

	v4 = 50

	// 匿名变量
	_, i, j, _ := 1, 2, 3, 4

	fmt.Println(v1, v2, v3, v4)

	fmt.Println("i =", i)
	fmt.Println("j =", j)

	// 匿名变量
	_, str1 := test()
	
	fmt.Println("str =", str1)

}
