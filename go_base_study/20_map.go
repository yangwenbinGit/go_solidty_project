package main

import "fmt"

func deleteMap(m map[int]string, key int) {
	delete(m, key)
}

func main() {
	var m1 map[int]string
	fmt.Println(m1)
	fmt.Println(m1 == nil)

	// 其他初始化方法
	m2 := map[int]string{}
	m3 := make(map[int]string)
	m4 := make(map[int]string, 10)
	fmt.Println(m2)
	fmt.Println(m3)
	fmt.Println(m4)

	// 定义的同时进行初始化操作
	m5 := map[int]string{1: "mike", 2: "yoyo", 3: "zhaoliu"}
	println(m5[1])
	println(m5[2])
	println(m5[3])
	println(m5[4]) // 空字符串 如果key不存在  会返回空字符串

	// 进行map的修改
	m5[1] = "aaaa"
	m5[3] = "bbbb"
	m5[4] = "cccc"

	fmt.Println(m5)

	// 往map中添加元素
	m6 := make(map[int]string, 10)
	m6[0] = "aaaaa"
	m6[1] = "bbbbb"
	m6[2] = "ccccc"
	m6[3] = "ddddd"
	m6[4] = "eeeee"
	fmt.Println(m6)

	// 遍历map数据
	for index, value := range m6 {
		fmt.Printf("key:%d,value:%s\n", index, value)
	}
	// 判断某个key是否存在
	value, ok := m6[1]
	if ok {
		fmt.Println("key存在,value为:", value)
	} else {
		fmt.Println("key不存在")
	}

	// 删除map中的元素
	delete(m6, 2)
	fmt.Println("删除后的map为:", m6)

	// 删除map中的元素(map作为函数参数也是引用传递)
	deleteMap(m6, 3)
	fmt.Println("调用函数deleteMap,删除后的map为:", m6)
}
