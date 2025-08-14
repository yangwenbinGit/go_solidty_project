package main

import "fmt"

type person01 struct {
	name string
	sex  byte
	age  int
}

type student struct {
	person01
	id  int
	add string
}

func (p *person01) printPerson01() {
	fmt.Printf("person01 name:%s, sex:%c, age:%d\n", p.name, p.sex, p.age)
}

// TODO 这个就叫做方法的重写 会覆盖上面的方法
func (s *student) printPerson01() {
	fmt.Printf("student name:%s, sex:%c, age:%d\n", s.name, s.sex, s.age)
}

func main() {
	// TODO 匿名字段 方法的集成 如果匿名字段实现了方法 那么包含这个匿名字段的结构体就可以直接调用方法

	p1 := person01{name: "张三", sex: 'm', age: 18}
	p1.printPerson01()

	// 定义一个学生 由于他里面包含了一个匿名字段 person01 所以他可以直接调用 person01 里面的方法
	s1 := student{
		person01: person01{
			name: "李四",
			sex:  'F',
			age:  18,
		},
		id:  1001,
		add: "北京",
	}
	//TODO 如果 student 里面实现了 printPerson01 方法 那么就会调用 student 里面的方法
	// 如果 student 里面没有实现 printPerson01 方法 那么就会调用 person01 里面的方法
	s1.printPerson01()

}
