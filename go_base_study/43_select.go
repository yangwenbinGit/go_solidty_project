package main

import "fmt"

// select
func fib(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			y, x = x, y+x
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			// 从通道中接收数据
			fmt.Println(<-c)
		}
		// 发送数据到通道
		quit <- 0
	}()

	fib(c, quit)

}
