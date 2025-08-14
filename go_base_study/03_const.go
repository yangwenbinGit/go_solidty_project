package main

func main() {
	const Pi float32 = 3.1415926

	const zero = 0.0

	const (
		size int64 = 1024
		eof        = 5
	)

	const name, gendery string = "张三", "男"

	const a, b, c = "1", "2", "3"

	// 如果用const定义的常量不可修改值

	println(a, b, c)
	println(size, eof)
	println(name, gendery)
}
