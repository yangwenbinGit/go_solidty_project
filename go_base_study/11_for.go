package main

import "fmt"

func main() {
	// 计算1-100的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("1-100的和为:", sum)
	fmt.Println("======================================================================")

	// range 关键字
	// 遍历数组  range 后面跟着要遍历的数组 返回两个值 第一个是索引  第二个是真正的值
	arr := []int{1, 2, 3, 4, 5}
	for index, value := range arr {
		fmt.Printf("index:%d value:%d \n", index, value)
	}
	fmt.Println("======================================================================")

	// continue 关键字
	// 跳过当前循环 继续下一次循环
	for i := 0; i <= 10; i++ {
		if i == 5 {
			continue
		} else {
			fmt.Printf("i:%d \n", i)
		}
	}

	fmt.Println("======================================================================")
	// break 关键字
	// 跳出当前循环
	for i := 0; i <= 10; i++ {
		if i == 9 {
			break
		} else {
			fmt.Printf("i:%d \n", i)
		}
	}

	fmt.Println("======================================================================")

	// goto
	for i := 1; i <= 5; i++ {
		for {
			fmt.Printf("i:%d \n", i)

			// 如果不加这句的话 内层就是一个死循环   goto跳转到哪里执行  执行完成后直接跳出整个外层循环
			goto MYGOTO
		}
		fmt.Printf("i:%d \n", i)
	}

MYGOTO:
	fmt.Println("这是goto语句")
}
