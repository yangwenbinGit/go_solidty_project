package main

import (
	"fmt"
	"time"
)

func main() {
	//创建定时器，每隔1秒后，定时器就会给channel发送一个事件(当前时间)
	ticker := time.NewTicker(time.Second * 1)

	i := 0
	go func() {
		for { //循环
			<-ticker.C
			i++
			fmt.Println("i = ", i)
			if i == 5 {
				ticker.Stop() //停止定时器
			}
		}
	}() //别忘了()

	time.Sleep(time.Second * 6)
}
