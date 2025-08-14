package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(time.Second * 2)
	t1 := time.Now() //当前时间
	fmt.Printf("t1: %v\n", t1)

	//TODO 从这个通道中读取数据 会阻塞 直到有数据可读
	t2 := <-timer1.C
	fmt.Printf("t2: %v\n", t2)

	// 如果只是单纯的等待的话
	timer2 := time.NewTimer(time.Second * 2)
	<-timer2.C
	fmt.Println("timer2 等待2s完成")

	time.Sleep(time.Second * 2)
	fmt.Println("再一次2s后")

	<-time.After(time.Second * 2)
	fmt.Println("再一次2s后")

	time3 := time.NewTimer(time.Second * 2)
	go func() {
		<-time3.C
		fmt.Println("time3 等待2s完成")
	}()

	isStop := time3.Stop()
	if isStop {
		fmt.Println("time3 停止成功") // 停止成功后 之前那个定时器就不会再执行了  time3 等待2s完成也没有打印
	}

}
