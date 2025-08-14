package main

import "fmt"

func main() {
	//1. 类型转换
	var ch byte = 'a'
	num := int(ch)
	println("num的值为:", num)

	var v1 int = 23
	v2 := float64(v1)
	fmt.Printf("v2的值为:%f", v2)

	// 2.类型别名
	type bigint int64
	var i bigint = 100
	fmt.Printf("i的值为:%d\n", i)

	type (
		MyInt    int64
		MyFloat  float64
		MyString string
	)

	var str MyString = "hello go"
	fmt.Printf("str的值为:%s\n", str)
}
