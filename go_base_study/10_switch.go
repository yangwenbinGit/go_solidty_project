package main

import "fmt"

func main() {

	// 第一种写法
	score := 90
	switch score {
	case 90:
		fmt.Println("优秀")
	case 80:
		fmt.Println("良好")
	case 70:
		fmt.Println("中等")
	case 60:
		fmt.Println("及格")
	default:
		fmt.Println("不及格")
	}

	// 第二种写法  直接声明
	switch score1 := 60; score1 {
	case 90:
		fmt.Println("优秀")
	case 80:
		fmt.Println("良好")
	case 70:
		fmt.Println("中等")
	case 60:
		fmt.Println("及格")
	default:
		fmt.Println("不及格")
	}

	// 第三种写法 可以写表达式 注意如果写表达式 switch的后面不能跟任何变量
	switch {
	case score >= 90:
		fmt.Println("优秀")
	case score >= 80:
		fmt.Println("良好")
	case score >= 70:
		fmt.Println("中等")
	case score >= 60:
		fmt.Println("及格")
	default:
		fmt.Println("不及格")
	}

}
