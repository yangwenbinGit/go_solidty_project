package main

import (
	"fmt"
	"strconv"
	"strings"
)

// strings 工具类的包集合操作
func main() {
	//1. contains 判断字符串是否包含指定的子字符串
	// 包含返回true 不包含返回false
	fmt.Println(strings.Contains("hello", "ll"))   // true
	fmt.Println(strings.Contains("wenbin", "bin")) // true

	// 2.join
	arr := []string{"hello", "world", "go"}
	str := strings.Join(arr, ",")
	fmt.Println(str) // hello,world,go

	// 3.Index 查找子字符串在字符串中的位置
	fmt.Println(strings.Index("hello", "ll"))   // 2
	fmt.Println(strings.Index("wenbin", "bin")) // 3

	// 4.Repeat 重复字符串
	fmt.Println("ba" + strings.Repeat("na", 2)) // banana

	// 5.Replace 替换字符串 在s字符串中，把old字符串替换为new字符串，n表示替换的次数，小于0表示全部替换
	fmt.Println(strings.Replace("oink oink oink", "k", "oo", 2))
	fmt.Println(strings.Replace("oink oink oink", "k", "dd", -1))

	// 6.split 字符串分割
	fmt.Println(strings.Split("hello,world,go", ","))
	fmt.Println(strings.Split("hello wenbin go", " "))

	// 7.Trim 在s字符串的头部和尾部去除cutset指定的字符串
	fmt.Println(strings.Trim("!!!     hello world  !", "! "))

	// 8.Fields 去除s字符串的空格符，并且按照空格分割返回slice
	fmt.Println(strings.Fields("foo bar baz")) // [foo bar baz]

	// 9.Append
	slice := make([]byte, 0, 100)
	slice = strconv.AppendInt(slice, 45679, 10)
	slice = strconv.AppendBool(slice, true)
	slice = strconv.AppendFloat(slice, 45.79, 'f', 2, 64)
	slice = strconv.AppendQuote(slice, "hello")
	fmt.Println(slice)

	// 10.Format
	a := strconv.FormatInt(123456, 10)
	b := strconv.FormatBool(true)
	c := strconv.FormatFloat(45.79, 'f', 2, 64)
	fmt.Println(a, b, c)

	aparse, err := strconv.ParseBool("false")
	CheckError(err)
	aparseInt, err := strconv.ParseInt("1234", 10, 64)
	aparseFloat, err := strconv.ParseFloat("56.89", 64)

	fmt.Println(aparse, aparseInt, aparseFloat)
}

func CheckError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
