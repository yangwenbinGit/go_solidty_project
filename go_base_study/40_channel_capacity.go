package main

import (
	"fmt"
	"time"
)

// 有缓存区的channel make(chan Type, capacity) capacity >0  表示缓存区的大小
func main() {
	c := make(chan int, 10)

	fmt.Println("channel的容量为:", cap(c))
	fmt.Println("channel的长度为:", len(c))

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
			fmt.Printf("子协程发送数据到channel中 i = %d, channel的长度为: %d, channel的容量为: %d\n", i, len(c), cap(c))
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 10; i++ {
		num := <-c
		fmt.Printf("主协程从channel中接收数据 num = %d, channel的长度为: %d, channel的容量为: %d\n", num, len(c), cap(c))
	}

	fmt.Println("主协程执行完成！")

	// 子线程和主线程的生产和消费他们是异步的
}
