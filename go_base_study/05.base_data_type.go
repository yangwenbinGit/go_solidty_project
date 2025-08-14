package main

import "fmt"

// 基本数据类型的练习

func main() {

	// 1.bool类型
	var v1 bool
	v1 = true

	v2 := false

	v3 := (1 == 2)

	var v4 = true

	fmt.Println(v1, v2, v3, v4)

	// 2.整形
	var v5 int
	v5 = 10

	var v6 = 20

	v7 := 30

	// cannot use -100 (untyped int constant) as uint value in variable declaration (overflows)
	// unit 默认是32 表示从0到42亿
	var v8 uint = 200

	fmt.Printf("v5 = %d, v6 = %d, v7 = %d, v8 = %d\n", v5, v6, v7, v8)

	// 3.浮点型
	var f1 float32
	f1 = 12.0

	f2 := 18.567

	fmt.Printf("f1 = %f, f2 = %f\n", f1, f2)

	// 4.字符byte

	var cha1, cha2, cha3 byte
	cha1 = 'a'
	cha2 = 'b'
	cha3 = 97
	fmt.Printf("char1 = %c, char2 = %c, char3 = %c\n", cha1, cha2, cha3)

	// 5.字符串
	var str1 string
	str1 = "hello world"

	str2 := "yangwenbin"

	var str3 = "hello go"

	fmt.Printf("str1 = %s, str2 = %s, str3 = %s\n", str1, str2, str3)

	// 6. 复数类型
	var c1 complex64
	c1 = 3.2 + 12i

	c2 := 3.2 + 12i // 默认是complex128 分为实部和虚部

	c3 := complex(3.2, 12) // 实部和虚部都为float64

	fmt.Printf("c1 = %v, c2 = %v, c3 = %v\n", c1, c2, c3)

	// 获取实部和虚部
	realC1 := real(c1)
	imagC1 := imag(c1)

	fmt.Printf("c1的实部为%v, c1的虚部为%v\n", realC1, imagC1)
}
