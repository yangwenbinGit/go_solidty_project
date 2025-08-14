package main

import "fmt"

// close() 关闭channel
func main() {

	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
			fmt.Println("子协程发送数据到channel中 i = ", i)
		}
		// 关闭channel  代表已经发送完成了
		close(c)
	}()

	for num := range c {
		// 这边不知道啥时候接收完成了  所以close(c) 关闭channel 代表已经接收完成了
		fmt.Println("主协程从channel中接收数据 num = ", num)
	}

	fmt.Println("接收数据完成!")

	// TODO 如果不写的话  会报错
	//	fatal error: all goroutines are asleep - deadlock!
	//
	//		goroutine 1 [chan receive]:
	//	main.main()
	//E:/MataNode/go_solidty_project/go_base_study/40_channel_close.go:18 +0xdf

}
