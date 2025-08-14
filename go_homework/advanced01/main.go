package main

import "fmt"

/**
136. 只出现一次的数字

给你一个 非空 整数数组 nums ，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
你必须设计并实现线性时间复杂度的算法来解决此问题，且该算法只使用常量额外空间。

示例 1 ：
输入：nums = [2,2,1]
输出：1

示例 2 ：
输入：nums = [4,1,2,1,2]
输出：4

示例 3 ：
输入：nums = [1]
输出：1


*/

func singleNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// 定义一个变量 用来保存结果
	res := 0
	mapResult := make(map[int]int)
	// 遍历数组
	for _, v := range nums {
		// 查看map中是否存在这个元素
		_, ok := mapResult[v]
		if ok {
			mapResult[v]++
		} else {
			mapResult[v] = 1
		}
	}
	// 遍历map
	for k, v := range mapResult {
		if v == 1 {
			res = k
		}
	}
	return res

}

func main() {
	nums := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(nums))

}
