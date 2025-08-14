package main

import "fmt"

// chan
/**
channel <- value 	// 发送value到channel
value := <- channel 	// 从channel接收值 并赋值给value
<- channel  接收并且丢弃数据
x, ok := <- channel 	// 从channel接收值 并赋值给value  如果channel关闭了  ok为false
*/

func main() {
	c := make(chan int)

	// 子协程执行
	go func() {
		defer fmt.Println("子协程执行完成")

		fmt.Println("子协程正在运行。。。")

		// 发送数据到channel
		c <- 100
	}()

	// 主协程执行
	num := <-c

	fmt.Println("num = ", num)
	fmt.Println("main协程结束")
}
