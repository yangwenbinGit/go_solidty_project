package main

import "fmt"

// panic 是一个内建函数，它停止当前 goroutine 的正常执行。
// 当函数 F 调用 panic，F 的正常执行就会停止，但是 F 中的延迟函数会正常执行，然后 F 返回到其调用者 G。
// 对 G 来说，F 就像调用了 panic 一样。
// 这个过程会持续下去，直到当前 goroutine 中所有的函数都返回，此时程序会崩溃。

func testA() {
	fmt.Println("func testA()")
}

func testB() {
	// 发生重大异常才用这个panic 直接会导致程序崩溃
	panic("func testB()")
	fmt.Println("func testB() 后面的代码")
}

func testC() {
	fmt.Println("func testC()")
}

func testD(index int) {
	arr := [10]int{}
	arr[index] = 100
	fmt.Println(arr)
}

func main() {
	testA()
	//testB() // 这里会导致程序崩溃 后面的代码不会执行
	testC()

	testD(20) // 这里数组越界 会导致程序崩溃

	/**
	func testA()
	panic: func testB()

	goroutine 1 [running]:
	main.testB(...)
		E:/MataNode/goCode/01_init_project/28_panic.go:16
	main.main()
		E:/MataNode/goCode/01_init_project/28_panic.go:25 +0x5b

	*/
}
