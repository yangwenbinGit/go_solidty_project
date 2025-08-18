package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// “这个操作不可分割，要么全做，要么不做” sync/atomic
func main() {

	var counter int64 // 须是 int64 等原子支持类型

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		// 每个协程对计数器进行+1操作
		wg.Add(1)
		go func() {
			// 执行完成后进行-1操作
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
		}()
	}

	// 等待所有协程执行完成 打印结果
	wg.Wait()
	fmt.Println("最终值:", counter) // 10

	//sync.Mutex：“锁住代码块，一次只能一个人进” → 适合复杂操作
	//sync/atomic：“操作本身不可分割” → 适合简单变量读写
}
