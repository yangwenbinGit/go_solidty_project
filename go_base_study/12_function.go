package main

import "fmt"

// 定义无参无返回值的函数
func test01() {
	fmt.Println("这是无参无返回值的函数test")
}

// 定义有参数无返回值的函数
func test1(a int) {
	fmt.Println("这是有参数无返回值的函数test1,参数为:", a)
}

// 定义有参有返回值的函数
func test2(a int, b int) (c int, d int) {
	//return a + b, a-b
	c = a + b
	d = a - b
	return
}

// 不定参数
func function01(args ...int) {
	fmt.Println("这是function01函数")
	for index, value := range args {
		fmt.Printf("index:%d value:%d \n", index, value)
	}
}

func function02(args ...int) {
	fmt.Println("这是function02函数")
	for index, value := range args {
		fmt.Printf("index:%d value:%d \n", index, value)
	}
}

func test03() (value int) {
	value = 100
	return
}

func main() {
	test01()
	test1(10)

	// 第一种方式
	//m, n := test2(10, 5)
	//println(m, n)

	// 第二种方式
	var m, n int
	m, n = test2(10, 5)
	fmt.Printf("m=%d, n=%d", m, n)

	// 调用function01函数
	function01(1, 2, 3, 4, 5)
	function02(30, 40, 50)

	a := test03()
	fmt.Printf("a的值为:%d", a)
}
