package main

import "fmt"

/**

1. 两数之和

给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。
你可以按任意顺序返回答案。

示例 1：

输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
示例 2：

输入：nums = [3,2,4], target = 6
输出：[1,2]
示例 3：

输入：nums = [3,3], target = 6
输出：[0,1]
*/

func twoSum(nums []int, target int) []int {
	// 定义一个map
	m := make(map[int]int)
	// 遍历数组
	for i := 0; i < len(nums); i++ {
		// 计算目标值与当前值的差值
		complement := target - nums[i]
		// 查看map中是否存在这个差值
		if _, ok := m[complement]; ok {
			// 如果存在，返回差值的索引和当前索引
			return []int{m[complement], i}
		}
		// 如果不存在，将当前值和索引加入map
		m[nums[i]] = i
	}
	// 如果没有找到，返回空数组
	return []int{}
}

func main() {
	//nums := []int{2, 7, 11, 15}
	nums := []int{3, 2, 4}
	target := 7
	result := twoSum(nums, target)
	fmt.Println(result) // [0 2]
}
