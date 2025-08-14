package main

import "fmt"

// 从控制台输入数据
func main() {
	var v int
	fmt.Println("请输入一个整数:")
	fmt.Scanf("%d", &v)
	fmt.Println("您输入的数字为:", v)
}
