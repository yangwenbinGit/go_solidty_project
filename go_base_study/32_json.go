package main

import (
	"encoding/json"
	"fmt"
)

type it struct {
	Company  string
	Subjects []string
	IsOk     bool
	Price    float64
}

// 可以定义输出的格式和方式
type it2 struct {
	// 如果是- 表示这个字段不会导出到json中
	Company string `json:"-"`
	//Subjects 的值会进行二次JSON编码
	Subjects []string `json:"subjects"`
	// 转换为字符串输出
	IsOk bool `json:",string"`
	// 转换为小写输出
	IsTrue bool `json:"isTrue"`
	// 如果 price 为 0 那么就不会输出
	Price float64 `json:",price,omitempty"`
}

func main() {
	// 1. 定义一个结构体
	it1 := it{
		Company:  "itcast",
		Subjects: []string{"go", "java", "python"},
		IsOk:     true,
		Price:    666,
	}
	// 将it1 转成 json 字符串
	jsonStr, err := json.Marshal(it1)
	if err != nil {
		fmt.Println("json.Marshal err:", err)
	} else {
		fmt.Println(string(jsonStr))
	}
	// {"Company":"itcast","Subjects":["go","java","python"],"IsOk":true,"Price":666}

	// 使用另外一种方式 就是格式化的方式 比较好看
	// 第一个参数是要转的结构体变量
	// 第二个参数是前缀 一般为空
	// 第三个参数是缩进 一般是两个空格
	jsonStr2, err := json.MarshalIndent(it1, "", "  ")
	if err != nil {
		fmt.Println("json.MarshalIndent err:", err)
	} else {
		fmt.Println(string(jsonStr2))
	}
	//{
	//	"Company": "itcast",
	//	"Subjects": [
	//	"go",
	//	"java",
	//	"python"
	//	],
	//	"IsOk": true,
	//	"Price": 666
	//}

	it2Test := it2{
		Company:  "itcast",
		Subjects: []string{"go", "java", "python"},
		IsOk:     true,
		IsTrue:   true,
		Price:    0,
	}
	jsonStr3, err := json.MarshalIndent(it2Test, "", "  ")
	if err != nil {
		fmt.Println("json.MarshalIndent err:", err)
	} else {
		fmt.Println(string(jsonStr3))
	}

	//{
	//	"subjects": [
	//	"go",
	//	"java",
	//	"python"
	//	],
	//	"IsOk": "true",
	//	"isTrue": true
	//}

	// 使用map转换为json 如果map 里面的value 是interface{} 那么就可以存储任意类型
	// 打印的时候是无序的
	map1 := make(map[string]interface{})
	map1["name"] = "wenbin"
	map1["age"] = 18
	map1["isOk"] = true
	map1["price"] = 666
	jsonStr4, err := json.MarshalIndent(map1, "", "  ")
	if err != nil {
		fmt.Println("json.MarshalIndent err:", err)
	} else {
		fmt.Println(string(jsonStr4))
	}

	//{
	//	"age": 18,
	//	"isOk": true,
	//	"name": "wenbin",
	//	"price": 666
	//}

}
