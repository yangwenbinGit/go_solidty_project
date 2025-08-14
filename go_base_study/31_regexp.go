package main

import (
	"fmt"
	"regexp"
)

// 正则表达式
func main() {
	context1 := "3.14 123123 .68 haha 1.0 abc 6.66 123."

	// 匹配的是 数字+小数点+数字
	reg := regexp.MustCompile(`\d+\.\d+`)
	// 查找所有的匹配
	findAll := reg.FindAllString(context1, -1)         // [3.14 1.0 6.66]
	result1 := reg.FindAllStringSubmatch(context1, -1) // [[3.14] [1.0] [6.66]]
	fmt.Println(findAll, result1)

}
