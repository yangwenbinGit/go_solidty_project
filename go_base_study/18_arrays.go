package main

import "fmt"

func modify(arr [10]int) (result [10]int) {
	arr[0] = 200
	fmt.Println("modify arr:", arr)
	return arr
}

func modifyPointer(arr *[5]int) {
	(*arr)[0] = 300
}

func main() {
	// 定义一个数组
	var arr [10]int

	for i := 0; i < 10; i++ {
		arr[i] = i
		fmt.Println("a[", i, "] =", arr[i])
	}

	// 遍历数组
	for index, value := range arr {
		fmt.Println("index的值为:", index, "value的值为:", value)
	}

	fmt.Println("数组的长度为:", len(arr))
	fmt.Println("数组的容量为:", cap(arr))

	// 初始化数组
	var arr1 [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(arr1)

	// 修改之前数组(要记住数组是值传递  在函数中修改完成后并不会更改原来的数组中的内容)
	fmt.Println("修改之前数组:", arr1)
	result := modify(arr1)
	fmt.Println("修改之后数组:", arr1)
	fmt.Println("接收修改后的数组:", result) // 接收后是修改后的数组 原来的数组没有改变

	// 传递一个指针数组
	var arr2 [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("修改之前数组:", arr2)
	modifyPointer(&arr2)
	fmt.Println("修改之后数组:", arr2)
}
