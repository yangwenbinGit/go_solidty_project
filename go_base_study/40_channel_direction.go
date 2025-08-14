package main

import "fmt"

//chan<- 表示数据进入管道，要把数据写进管道，对于调用者就是输出。
//<-chan 表示数据从管道出来，对于调用者就是得到管道的数据，当然就是输入。

func counter(in chan<- int) {
	for i := 0; i < 10; i++ {
		in <- i
	}
	defer close(in)
}

func printData(out <-chan int) {
	for num := range out {
		fmt.Println(num)
	}
	fmt.Println("读取完成")
}

func main() {

	//TODO 无方向的通道 可以发送数据 也可以接收数据
	//c0 := make(chan int, 10)
	//// 通道的方向
	////TODO  chan<- 表示数据进入管道，要把数据写进管道，对于调用者就是输出。
	//c1 := make(chan<- int)
	//
	//// TODO <-chan 表示数据从管道出来，对于调用者就是得到管道的数据，当然就是输入。
	//c2 := make(<-chan int)

	//fmt.Println(c1)
	//fmt.Println(c2)
	//fmt.Println(c0)

	c := make(chan int)

	// 启动一个协程 向通道中发送数据
	go counter(c)
	printData(c)
}
