package main

import "fmt"

/*
*
✅指针
题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
考察点 ：指针的使用、值传递与引用传递的区别。
题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
考察点 ：指针运算、切片操作。
*/

// 指针
func addPointer(number *int) int {
	if number == nil {
		return 0
	}
	*number += 10
	return *number
}

// 切片
func addCut(number *[]int) []int {
	if len(*number) == 0 {
		return *number
	}
	for i := 0; i < len(*number); i++ {
		(*number)[i] *= 2
	}
	return *number
}

func main() {

	var i int = 10
	addresult := addPointer(&i)
	fmt.Println("调用函数addPointer的结果为:", addresult) // 20

	var number []int = []int{1, 2, 3, 4, 5}
	addCutResult := addCut(&number)
	fmt.Println("调用函数addCut的结果为:", addCutResult) // [2 4 6 8 10]
}
