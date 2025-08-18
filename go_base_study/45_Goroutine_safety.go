package main

import (
	"fmt"
	"time"
)

// go语言中多个协程对同一个变量进行操作时，可能会出现竞态条件，导致最终结果不确定 最终的结果和我们想要的结果不一样的
func main() {
	var counter int

	// 启动10个协程  每一个协程对计数器进行1000次递增操作
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				counter++
			}
		}()
	}

	time.Sleep(2 * time.Second)  // 等待
	fmt.Println("最终值:", counter) // 期望 10000，但结果可能是 最终值: 9341  最终值: 6756
}
