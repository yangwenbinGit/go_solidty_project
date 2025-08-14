package main

import (
	"fmt"
	"sort"
)

/**
56. 合并区间

以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。

示例 1：
输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].


示例 2：
输入：intervals = [[1,4],[4,5]]
输出：[[1,5]]
解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。
*/

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}
	// 先排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 定义一个结果数组
	res := make([][]int, 0)
	// 定义一个临时数组
	temp := intervals[0]
	// 遍历数组
	for i := 1; i < len(intervals); i++ {
		// 如果当前数组的第一个元素大于临时数组的第二个元素
		if intervals[i][0] > temp[1] {
			// 说明没有重叠
			// 把临时数组加入结果数组
			res = append(res, temp)
			// 把当前数组赋值给临时数组
			temp = intervals[i]
		} else {
			// 说明有重叠
			// 比较当前数组的第二个元素和临时数组的第二个元素
			if intervals[i][1] > temp[1] {
				// 说明当前数组的第二个元素更大
				// 把当前数组的第二个元素赋值给临时数组的第二个元素
				temp[1] = intervals[i][1]
			}
		}
	}
	// 把临时数组加入结果数组
	res = append(res, temp)

	return res
}

func main() {
	intervals := [][]int{{1, 2}, {2, 4}, {4, 5}}
	result := merge(intervals)
	fmt.Println(result) // [[1 5]]
}
