package main

import "fmt"

func testSlice(s []int) { //切片做函数参数
	s[0] = -1
	fmt.Println("testSlice : ")
	for index, value := range s {
		fmt.Printf("%d,%d\n", index, value)
	}
}

// 切片
func main() {
	// 定义一个切片 切片和数组的区别是 数组的长度是固定的 切片的长度不写
	var s1 []int
	s2 := []int{}
	s3 := []int{1, 2, 3}
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)

	// 使用make创建切片 make一共有三个参数 第一个参数是类型 第二个参数是长度 第三个参数是容量  第三个不写和长度一样的
	var s4 []int = make([]int, 10, 20)
	fmt.Println(s4)
	fmt.Println(len(s4))
	fmt.Println(cap(s4))

	// make只能创建slice、map和channel，并且返回一个有初始值(非零)。
	s5 := make([]int, 10)
	fmt.Println(s5)

	// 切片对应的操作
	var s6 []int = make([]int, 10)
	s6 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(s6)

	fmt.Println("===============================================")
	fmt.Println(s6[2])
	fmt.Println(s6[:])
	fmt.Println(s6[2:])
	fmt.Println(s6[:8])  // 不包含最后一个  [1 2 3 4 5 6 7 8]
	fmt.Println(s6[2:5]) // [3 4 5]
	fmt.Println(s6[2:8:10])
	fmt.Println(len(s6)) // len = high -low
	fmt.Println(cap(s6)) // cap = max - low

	fmt.Println("===============================================")
	// 对于切片需要注意的是如果修改里面的内容  那会影响到原来的切片内容  内容是真正发生改变的
	fmt.Println("修改之前的切片", s6)
	s6[2] = 100
	fmt.Println("修改之后的切片", s6)

	s7 := s6[3:6]
	fmt.Println("修改之前的切片", s7)
	s7[0] = 200
	fmt.Println("修改之后的切片", s7)
	fmt.Println("修改之后的切片s6", s6)

	// append 向切片尾部追加元素
	s6 = append(s6, 3000, 4000, 5000)
	fmt.Println("append之后的切片s6", s6) // append之后的切片s6 [1 2 100 200 5 6 7 8 9 10 3000 4000 5000]
	//TODO 如果超过了容量  会自动扩容  扩容的规则是  容量*2

	testSlice(s6)
	fmt.Println("修改之后的切片s6", s6)
}
