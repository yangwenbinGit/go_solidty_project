package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

/**
锁机制
题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ： sync.Mutex 的使用、并发数据安全。
题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ：原子操作、并发数据安全。
*/

func mutexDemo() {
	var counter int64
	var mu sync.Mutex

	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				mu.Lock()
				defer mu.Unlock()
				counter++
			}
		}()
	}

	time.Sleep(2 * time.Second)

	fmt.Printf("%d\n", counter)
}

func atomicDemo() {
	var counter int64
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&counter, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Printf("%d\n", counter)
}

func mutexDemoFix() {
	var counter int64
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done() // 确保一定会调用
			for j := 0; j < 1000; j++ {
				mu.Lock()   // 获取锁
				counter++   // 修改共享变量
				mu.Unlock() // 立即释放锁（或用 defer）
			}
		}()
	}

	wg.Wait() // 等待所有协程完成
	fmt.Printf("counter = %d\n", counter)
}

func main() {
	mutexDemo() // 第一个慎用 有时候有问题
	mutexDemoFix()
	atomicDemo()
}
