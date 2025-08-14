package main

import "fmt"

func swap(a, b int) {
	a, b = b, a
	fmt.Println("swap result:", a, b)
}

func swap02(x, y *int) {
	*x, *y = *y, *x
	fmt.Println("swap result:", *x, *y)
}

// 指针
func main() {
	var a int = 100
	fmt.Printf("%p\n", &a)

	var p *int
	p = &a
	fmt.Printf("%p\n", p)
	fmt.Printf("a=%d, *p=%d\n", a, *p)

	// 给指针变量赋值
	*p = 200
	fmt.Printf("%p\n", p)
	fmt.Printf("a=%d, *p=%d\n", a, *p)

	// 创建对象
	var p1 *int = new(int)
	fmt.Printf("%p\n", p1)
	fmt.Printf("%d\n", *p1)

	var p2 *int = new(int)
	*p2 = 1000
	fmt.Printf("%p\n", p2)
	fmt.Printf("%d\n", *p2)

	// 调用普通函数
	var x, y int = 100, 200
	swap(x, y)
	fmt.Printf("x=%d, y=%d\n", x, y) // 这个地方特别注意 普通函数调用 不会改变参数的值 只是改变了函数内部的变量值

	// 调用指针参数的函数
	swap02(&x, &y)
	fmt.Printf("x=%d, y=%d\n", x, y) // 这个地方如果传的是指针 那么会真正改变参数的值 现在执行完成后 x真的变成200  y变成100

}
