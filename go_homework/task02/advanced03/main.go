package main

import "fmt"

/*
*题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
考察点 ：接口的定义与实现、面向对象编程风格。
题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
考察点 ：组合的使用、方法接收者。
*/

// 定义接口
type Shape interface {
	Area()
	Perimeter()
}

type Rectangle struct {
	name string
}

func (r *Rectangle) Area() {
	fmt.Println("Rectangle中实现Area打印结果为", r.name)
}

func (r *Rectangle) Perimeter() {
	fmt.Println("Rectangle中实现Perimeter打印结果为", r.name)
}

type Circle struct {
	name string
}

func (r *Circle) Area() {
	fmt.Println("Circle中实现Area打印结果为", r.name)
}

func (r *Circle) Perimeter() {
	fmt.Println("Circle中实现Perimeter打印结果为", r.name)
}

func commonShape(sh Shape) {
	sh.Area()
	sh.Perimeter()
}

// =======================================================================================================

type Person struct {
	name string
	age  int
}

type Employee struct {
	Person
	EmployeeID string
}

func (emp *Employee) PrintInfo() {
	fmt.Printf("员工的姓名为:%s,年龄为:%d,员工ID为:%s", emp.name, emp.age, emp.EmployeeID)
}

func main() {
	// 调用
	re := Rectangle{"长方形"}
	ci := Circle{"圆形"}

	commonShape(&re)
	commonShape(&ci)
	/**
	Rectangle中实现Area打印结果为 长方形
	Rectangle中实现Perimeter打印结果为 长方形
	Circle中实现Area打印结果为 圆形
	Circle中实现Perimeter打印结果为 圆形
	*/
	// 调用
	emp := Employee{
		Person: Person{
			name: "张三",
			age:  18,
		},
		EmployeeID: "001",
	}
	emp.PrintInfo()
	/**
	员工的姓名为:张三,年龄为:18,员工ID为:001
	*/

}
