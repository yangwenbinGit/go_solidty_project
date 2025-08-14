package main

import (
	"fmt"
	"time"
)

func newTask() {
	var i int
	for i = 1; i <= 10; i++ {
		fmt.Println("newTask任务执行:", i)
		time.Sleep(1 * time.Second) // 延时1秒
	}
}

func main() {

	// 启动一个newTask任务 两个任务交替执行
	go newTask()

	//var i int
	//for i = 1; i <= 10; i++ {
	//	fmt.Println("main任务执行:", i)
	//	time.Sleep(1 * time.Second) // 延时1秒
	//}

	// TODO 这里要特别的注意  newTask任务是异步执行的 所以main任务执行完成 不会等待newTask任务执行完成
	// TODO 意思就是当main任务执行完成 直接就退出了  当main任务退出 newTask任务也会强制退出
	fmt.Println("main任务执行完成")

	//main任务执行完成
	//newTask任务执行: 1

}
