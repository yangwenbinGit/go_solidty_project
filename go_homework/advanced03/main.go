package main

import "fmt"

/**
给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。这些数字按从左到右，从最高位到最低位排列。这个大整数不包含任何前导 0。

将大整数加 1，并返回结果的数字数组。

示例 1：

输入：digits = [1,2,3]
输出：[1,2,4]
解释：输入数组表示数字 123。
加 1 后得到 123 + 1 = 124。
因此，结果应该是 [1,2,4]。
示例 2：

输入：digits = [4,3,2,1]
输出：[4,3,2,2]
解释：输入数组表示数字 4321。
加 1 后得到 4321 + 1 = 4322。
因此，结果应该是 [4,3,2,2]。
示例 3：

输入：digits = [9]
输出：[1,0]
解释：输入数组表示数字 9。
加 1 得到了 9 + 1 = 10。
因此，结果应该是 [1,0]。
*/

func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return digits
	}
	// 将数组转换为整数
	num := 0
	for i := 0; i < len(digits); i++ {
		num = num*10 + digits[i]
	}
	// 然后执行+1
	num++
	// 然后将整数转为数组
	// 定义一个数组
	res := make([]int, 0)
	// 然后将整数转为数组
	for num > 0 {
		res = append(res, num%10)
		num = num / 10
	}
	// 反转数组
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}
	return res
}

func main() {
	//digits := []int{1, 2, 3}
	digits := []int{4, 3, 2, 1}
	fmt.Println(plusOne(digits))
}
