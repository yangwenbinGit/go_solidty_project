package main

import "fmt"

type humaner01 interface {
	sayHi()
}

type person02 interface {
	humaner01 // 这里就和写了sayHi()方法一样的
	sing(songName string)
}

type student02 struct {
	id    int
	name  string
	score float32
}

// 实现接口方法sayHi
func (s *student02) sayHi() {
	fmt.Printf("我是:%s, 我是一个学生\n", s.name)
}

// 实现接口方法sing
func (s *student02) sing(songName string) {
	fmt.Printf("我是:%s, 我会唱歌:%s\n", s.name, songName)
}

type element interface{}

// 接口嵌入
func main() {
	// 定义一个学生
	s1 := &student02{1, "张三", 90}
	// 调用方法
	s1.sayHi()
	s1.sing("黄昏")

	//TODO 接口可以进行转换 超集可以转换为子集  反之会报错
	//TODO 在这个例子中 person02 是一个超集  humaner01 是一个子集

	var p1 person02 = s1
	var h1 humaner01 = p1 //TODO 可以将超集转换为子集
	h1.sayHi()
	p1.sing("有一天")

	// 还有一些空接口
	var v1 interface{} = 1
	var v2 interface{} = "abc"
	var v3 interface{} = true
	var v4 interface{} = s1
	fmt.Printf("v1的类型为:%T, v1的值为:%v\n", v1, v1)
	fmt.Printf("v2的类型为:%T, v2的值为:%v\n", v2, v2)
	fmt.Printf("v3的类型为:%T, v3的值为:%v\n", v3, v3)
	fmt.Printf("v4的类型为:%T, v4的值为:%v\n", v4, v4)

	fmt.Println("========================================================")

	//value, ok = element.(T)，这里value就是变量的值，ok是一个bool类型，element是interface变量，T是断言的类型
	list := make([]element, 5)
	list[0] = "hello"
	list[1] = 100
	list[2] = 3.45
	list[3] = &student02{1, "张三", 90}
	list[4] = true
	for index, v := range list {
		fmt.Printf("index:%d, v:%v\n", index, v)
		switch v.(type) {
		case string:
			fmt.Printf("v是一个字符串类型, 值为:%v\n", v)
		case int:
			fmt.Printf("v是一个int类型, 值为:%v\n", v)
		case float64:
			fmt.Printf("v是一个float32类型, 值为:%v\n", v)
		case *student02:
			fmt.Printf("v是一个*student02类型, 值为:%v\n", v)
		case bool:
			fmt.Printf("v是一个bool类型, 值为:%v\n", v)
		default:
			fmt.Printf("v的类型未知, 值为:%v\n", v)
		}
	}
}
