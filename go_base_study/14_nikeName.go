package main

import "fmt"

// 在go语言中所有的匿名函数都是闭包  匿名函数就是没有名字的函数

func main() {

	i := 10
	str := "hello go"
	// 方式1 匿名函数无参数 无返回值
	f1 := func() {
		fmt.Println("匿名函数无参数 无返回值", i, str)
	}

	// 调用函数
	f1()

	// 声明一个函数类型 名字是FuncType 然后将函数f1赋值给f2  直接调用f2() 相当于调用的f1()
	type FuncType func()
	var f2 FuncType = f1
	f2()

	// 第三中方式

	f3 := func() {
		fmt.Println("f3匿名函数无参数 无返回值", i, str)
	}
	f3()

	fmt.Println("=============================================================")
	// 第四中方式
	func() {
		i = 100
		str = "hello world"

		fmt.Printf("内部i=%d,str=%s\n", i, str)

	}()

	fmt.Printf("外部i=%d,str=%s\n", i, str)

}
