package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

type student03 struct {
	Person // 这里写了一个匿名字段 那么这个结构体就有了Person的所有字段
	id     int
	add    string
}

type student04 struct {
	Person // 这里写了一个匿名字段 那么这个结构体就有了Person的所有字段
	id     int
	add    string
	name   string
}

// 相当于给string起了一个别名
type mystr string

type student05 struct {
	Person // 匿名字段 结构体类型
	int    // 匿名字段  内置类型
	mystr  // 匿名字段 自定义类型
}

// 结构体指针类型
type student06 struct {
	*Person // 这里写了一个匿名字段 那么这个结构体就有了Person的所有字段
	id      int
	add     string
}

func main() {
	var s1 student03 = student03{
		Person: Person{
			name: "mike",
			sex:  'm',
			age:  18,
		},
		id:  1,
		add: "beijing",
	}
	fmt.Printf("%+v\n", s1)

	// 还有一种赋值方式
	s2 := student03{Person{"mike", 'm', 18}, 2, "shanghai"}
	fmt.Printf("%+v\n", s2)

	s3 := student03{Person{name: "mike", sex: 'm', age: 18}, 3, "shanghai"}
	fmt.Printf("%+v\n", s3)

	s4 := student03{Person{name: "mike", sex: 'm'}, 4, "chongqing"}
	fmt.Printf("%+v\n", s4)

	//成员的操作
	var s5 student03
	s5.id = 5
	s5.name = "mike"
	s5.sex = 'm'
	s5.age = 18
	s5.add = "beijing"
	fmt.Printf("%+v\n", s5)

	// 如果在student04也有和Person的字段名相同的字段 那么就会覆盖Person的字段
	//TODO 表明这里是最近原则 赋值的时候先去找student04的name 如果它里面没有  那么就会去Person里面找 然后进行赋值
	var s6 student04
	s6.name = "mike"
	fmt.Printf("%s\n", s6.Person.name) // 这里打印的是空
	fmt.Printf("%+v\n", s6)            // {Person:{name: sex:0 age:0} id:0 add: name:mike}

	// TODO 这里的int是匿名字段 所以这里的int是可以直接使用的 我自己定义的mystr自定义字段 也可以直接使用  s7.int  s7.mystr 就可以获取到值
	s7 := student05{Person{"mike", 'm', 18}, 100, "beijing"}
	fmt.Printf("name:%s, sex:%c, age:%d, id:%d, add:%s\n", s7.name, s7.sex, s7.age, s7.int, s7.mystr)

	// TODO 给结构体指针类型赋值
	s8 := student06{&Person{"mike", 'm', 18}, 3, "shanghai"}
	fmt.Println(*(s8.Person), s8.id, s8.add)

}
