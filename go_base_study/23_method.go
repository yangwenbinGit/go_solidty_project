package main

import "fmt"

// 自定义类型
type myInt int

// TODO 在函数定义的时候 在函数名的前面有一个变量 这个变量可以直接调用函数
func (a myInt) add(b myInt) myInt {
	return a + b
}

// 最开始的写法
func add01(a, b myInt) myInt {
	return a + b
}

// 结构体
type person struct {
	name string
	sex  byte
	age  int
}

// 结构体作为参数进行传递
func (p person) printInfo() {
	fmt.Printf(" name:%s, sex:%c, age:%d\n", p.name, p.sex, p.age)
}

func (p person) setInfoValue() {
	p.name = "张三"
	p.sex = 'm'
	p.age = 18
}

func (p *person) setInfoPointer() {
	p.name = "李四"
	p.sex = 'm'
	p.age = 18
}

func main() {
	// 第一种是通过前面变量进行调用
	var a myInt = 10
	var b myInt = 20
	var c myInt = a.add(b)
	fmt.Println("a.add(b)的结果是:", c)

	// 第二种是原始的写法
	var d myInt = add01(a, b)
	fmt.Println("add01(a, b)的结果是:", d)

	// 定义pserson
	p1 := person{"mike", 'm', 18}
	p1.printInfo()

	// 如果作为值传递的话 这个结构体的数值不会改变
	p2 := person{"mike", 'm', 18}
	fmt.Printf("p2的值为:%+v\n", p2)
	p2.setInfoValue()
	fmt.Printf("设置完后p2的值为:%+v\n", p2)

	// 如果作为指针传递的话 这个结构体的数值会改变
	p3 := person{"mike", 'm', 18}
	fmt.Printf("p3的值为:%+v\n", p3)
	//(&p3).setInfoPointer()
	p3.setInfoPointer() // 这个地方不加& 也可以  因为go语言会自动的转换
	fmt.Printf("设置完后p3的值为:%+v\n", p3)

	//TODO 对于setInfoValue 也可以给他传递指针 (他会自动转换 指针变为非指针操作 变为值传递)
	p4 := &person{"mike", 'm', 18}
	(p4).setInfoValue()
	(*p4).setInfoValue()
	fmt.Printf("设置完后p4的值为:%+v\n", p4)

}
