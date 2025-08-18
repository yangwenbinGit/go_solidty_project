package main

import (
	"fmt"
	"time"
)

/*
题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
考察点 ： go 关键字的使用、协程的并发执行。

*/

func printOdd(num int) {
	for i := 1; i <= num; i++ {
		if i%2 != 0 {
			fmt.Println("奇数:", i)
		}
	}
}

func printEven(num int) {
	for i := 1; i <= num; i++ {
		if i%2 == 0 {
			fmt.Println("偶数:", i)
		}
	}
}

func main() {

	go printOdd(10)

	go printEven(10)

	time.Sleep(2 * time.Second)

}
