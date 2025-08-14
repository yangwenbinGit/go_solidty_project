package main

import "fmt"

type humaner interface {
	sayHi()
}

type stu struct {
	name  string
	score float64
}

// 学生实现这个sayHi方法
func (s *stu) sayHi() {
	fmt.Printf("stu类型 我是%s, 我考了%f分\n", s.name, s.score)
}

type teacher struct {
	name  string
	group string
}

func (t *teacher) sayHi() {
	fmt.Printf("teacher类型 name:%s, group:%s\n", t.name, t.group)
}

type myStr string

func (s *myStr) sayHi() {
	fmt.Printf("myStr类型 我是%s\n", *s)
}

// 写一个公共的接口类 传递接口类型
// TODO 这里特别的注意 go语言不允许接口的类型是一个指针 这样会报错的 和你传递的变量有关系 如果是& 那么过来就是指针类型
func commonSayHi(h humaner) {
	h.sayHi()
}

func main() {

	//TODO 如果我们不用统一接口  各自实现
	// 学生
	s1 := stu{"张三", 90}
	s1.sayHi()
	// 老师
	t1 := teacher{"李四", "1班"}
	t1.sayHi()
	// 字符串
	s2 := myStr("hello")
	s2.sayHi()

	fmt.Println("=====================================================")
	// 调用同一接口，不同表现
	commonSayHi(&s1)
	commonSayHi(&t1)
	commonSayHi(&s2)

	fmt.Println("=====================================================")
	// 也可以弄一个切片
	x := make([]humaner, 3)
	x[0], x[1], x[2] = &s1, &t1, &s2
	for _, v := range x {
		v.sayHi()
	}
}
