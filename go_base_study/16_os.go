package main

import (
	"fmt"
)

var x int = 200

func test001(a, b int) {
	var c int = 100
	fmt.Println(a, b, c)

	fmt.Println("x的值为:", x)
}

func main() {
	//args := os.Args
	//
	//if args == nil || len(args) < 2 {
	//	fmt.Println("请输入参数")
	//	return
	//}
	//ip := args[1]
	//port := args[2]
	//fmt.Printf("ip:%s,port:%s", ip, port)

	// 在main函数中我们是无法直接访问变量 a b的 还有c变量 因为他们属于局部变量
	// 但是我们可以在main函数中调用test函数 因为test函数的参数是a b 而main函数中也有a b变量
	test001(10, 20)
	{
		var i int = 100
		fmt.Println(i)

	}

	fmt.Println("main函数打印x的值为:", x)

}
