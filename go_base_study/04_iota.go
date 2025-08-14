package main

// 在一个有const声明的语句中，第一个声明的iota的值为0  第二个为1  依次累加

func main() {

	// 在不同行的话 每次加1
	const (
		a = iota
		b = iota
		c = iota

		w
	)

	// 如果在同一行的话  他们的值都是一样的
	const (
		h, i, j = iota, iota, iota
	)

	const (
		k = iota * 10
		l = iota * 10
		m = iota * 10
	)

	println(a, b, c, w)
	println(h, i, j)
	println(k, l, m)

}
