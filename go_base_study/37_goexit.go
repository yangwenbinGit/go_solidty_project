package main

import (
	"fmt"
	"runtime"
	"time"
)

// goexit 用于退出当前的goroutine

func testA_01() {
	defer fmt.Println("testA_01 执行")
	testA_02()
}

func testA_02() {
	defer fmt.Println("testA_02 执行")
	runtime.Goexit()
	fmt.Println("testA_02 执行结束")
}

func main() {

	go testA_01()

	time.Sleep(1000 * time.Millisecond)
}

// defer 这个是先进后出 先看到的会最后执行   在testA_02中遇到了runtime.Goexit()会直接终止当前子线程的执行
//testA_02 执行
//testA_01 执行
