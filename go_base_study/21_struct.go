package main

import "fmt"

// 21_结构体的学习
type Student struct {
	id      int
	name    string
	sex     byte
	age     int
	address string
}

func printStudentValue(s Student) {
	s.name = "peter1"
}

func printStudentValuePointer(s *Student) {
	s.name = "peter1"
}

func main() {
	var s1 Student = Student{1, "mike", 'm', 18, "beijing"}
	fmt.Println(s1)

	s2 := Student{2, "tom", 'm', 20, "beijing"}
	fmt.Println(s2)

	s3 := Student{id: 1, name: "mike", sex: 'm', age: 20, address: "beijing"}
	fmt.Println(s3)

	s4 := Student{
		id:   1,
		name: "mike",
	}
	fmt.Println(s4)

	// 指针变量
	var s5 *Student = &Student{5, "wenbin", 'm', 30, "beijing"}
	fmt.Println(*s5) // 打印的值
	fmt.Println(&s5) // 打印的地址

	s6 := &Student{
		id:      6,
		name:    "mike",
		sex:     'm',
		age:     30,
		address: "shanghai",
	}
	fmt.Println(*s6)
	fmt.Println(&s6)

	// 打印成员
	fmt.Printf("s6的id是:%d, name是:%s, sex是:%c, age是:%d, address是:%s\n", s6.id, s6.name, s6.sex, s6.age, s6.address)

	//给成员变量赋值
	s6.id = 100
	s6.name = "peter"
	s6.sex = 'M'
	s6.address = "beijing"
	fmt.Printf("s6的id是:%d, name是:%s, sex是:%c, age是:%d, address是:%s\n", s6.id, s6.name, s6.sex, s6.age, s6.address)

	s7 := new(Student)
	s7.id = 200
	s7.name = "wenhao"
	s7.sex = 'F'
	fmt.Printf("s7的id是:%d, name是:%s, sex是:%c\n", s7.id, s7.name, s7.sex)

	// 普通类型和指针类型打印
	fmt.Printf("s7的值为:%v, 取地址符号为：%v\n", s7, &s7)

	var p1 *Student = s7
	p1.id = 300
	(*p1).name = "mike"
	(*p1).sex = 'M'
	(*p1).age = 30
	(*p1).address = "shanghai"
	fmt.Printf("p1的id是:%d, name是:%s, sex是:%c\n", p1.id, p1.name, p1.sex)

	// 结构体的比较
	var s8 Student = Student{1, "mike", 'm', 18, "beijing"}
	var s9 Student = Student{1, "mike", 'm', 18, "beijing"}
	fmt.Println("s8和s9是否相等:", s8 == s9)  // true
	fmt.Println("s8和s9是否不相等:", s8 != s9) // false

	// TODO 记住这个结构体  传递的是值类型  所以不会改变原来的值 只会在执行方法的内部改变 并不会改变原来的值
	printStudentValue(s8)
	fmt.Println(s8)

	// 如果是指针传递的话 会改变这个结构体原来的值
	printStudentValuePointer(&s9)
	fmt.Println(s9)

}

// TODO 在go语言中如果我们的这个结构体要被别的文件引用 首字母要大写
// 这个里面小写的话不能被别的文件引用  大写的02才可以
type student01 struct {
	id   int
	name string
}

type Student02 struct {
	id   int
	name string
}
