package main

import "fmt"

func main() {
	// 运算符 + =- * /
	fmt.Println(10 + 5)
	fmt.Println(10 - 5)
	fmt.Println(10 * 5)
	fmt.Println(10 / 5)
	fmt.Println(10 % 5)
	a := 0
	a++
	fmt.Println("a的值为:", a)

	b := 2
	b--
	fmt.Println("b的值为:", b)

	fmt.Println("===============================================================")
	// 关系运算符==   !=  >  <    <=  >=
	fmt.Println(4 == 3)
	fmt.Println(4 != 3)
	fmt.Println(4 > 3)
	fmt.Println(4 < 3)
	fmt.Println(4 <= 3)
	fmt.Println(4 >= 3)
	fmt.Println(4 >= 4)

	fmt.Println("===============================================================")

	// 逻辑运算符 ! 非  && 与  || 或
	fmt.Println(!true)
	fmt.Println(true && false)
	fmt.Println(true || false)

	fmt.Println("===============================================================")

	// 逻辑运算符 短路运算
	var a1 int = 10
	var b1 int = 20
	fmt.Println(&a1)
	fmt.Println(&b1)

	fmt.Println("===============================================================")

	var c1 *int
	c1 = &a1

	// 下面&c1 是指针  *c1 是指针指向的值  这个地方需要注意的是 &a1虽然复制给指针c1  c1指向的地址和a1的地址是一样的 但是c1也有自己的
	fmt.Println("c1指向的地址为:", &c1)
	fmt.Println("c1指向的值为:", *c1)

}
