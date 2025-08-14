package main

import (
	"fmt"
	"runtime"
	"time"
)

// runtime.Gosched() 用于让出CPU时间片，让出当前goroutine的执行权限，调度器安排其他等待的任务运行，并在下次某个时候从该位置恢复执行。

func newTask01(name string) {
	for i := 0; i < 3; i++ {
		fmt.Println(name)
	}
}

func newTask02() {
	for i := 0; i < 3; i++ {
		fmt.Println("new task02 任务执行")
	}
}

func main() {
	go newTask01("new task01")

	go newTask02()

	// 添加短暂等待 真正阻塞主线程 让其他的子线程执行完成 放在这里让其他线程先执行
	time.Sleep(100 * time.Millisecond)

	for i := 0; i < 3; i++ {
		runtime.Gosched()
		fmt.Println("main hello")
	}
}

//TODO 没有加 runtime.Gosched()的打印结果
//main hello
//new task01
//new task01
//new task01
//main hello
//main hello

// TODO 加了runtime.Gosched()之后 main函数在执行的时候遇到了runtime.Gosched()后就释放出cpu的执行权限,让其他的线程先执行
//new task01
//new task01
//new task01
//main hello
//main hello
//main hello
