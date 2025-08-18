package main

import (
	"fmt"
	"time"
)

/*
*✅Channel
题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
考察点 ：通道的基本使用、协程间通信。
题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
考察点 ：通道的缓冲机制
*/

// 生产者
func produce(ch chan int, num int) {
	for i := 1; i <= num; i++ {
		ch <- i
		fmt.Println("生产者发送数据:", i)
	}
}

// 消费者
func consume(ch chan int, num int) {
	for i := 0; i < num; i++ {
		fmt.Println("消费者接收数据:", <-ch)
	}
}

func main() {
	//ch := make(chan int)
	//go produce(ch)
	//go consume(ch)
	//// 等待生产者和消费者协程执行完毕 等待2秒
	//time.Sleep(2 * time.Second)

	ch := make(chan int, 100)
	go produce(ch, 100)
	go consume(ch, 100)

	// 等待生产者和消费者协程执行完毕 等待2秒
	time.Sleep(2 * time.Second)

}
