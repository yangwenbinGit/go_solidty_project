package main

import (
	"fmt"
	"sync"
	"time"
)

// go语言中多个协程对同一个变量进行操作时，可能会出现竞态条件，导致最终结果不确定 最终的结果和我们想要的结果不一样的
// 针对这个问题 我们可以使用锁机制来解决  sync.Mutex：互斥锁  TODO ：一次只能一个人访问共享资源
// 适用场景：多个 Goroutine 要修改同一块数据（如 map、slice、变量）
func main() {
	var counter int
	var mu sync.Mutex

	// 启动10个协程  每一个协程对计数器进行1000次递增操作
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				mu.Lock() // 加锁：只有我能改 别人就要等着
				// TODO： 这个地方需要注意 解锁加上defer 确保可以解锁 如果业务发生异常 也可以解锁
				defer mu.Unlock()
				counter++
				//mu.Unlock() // 解锁：我改完了  其他的人可以改了

			}
		}()
	}

	time.Sleep(2 * time.Second)  // 等待
	fmt.Println("最终值:", counter) // 期望 10000 实际 10000
}
