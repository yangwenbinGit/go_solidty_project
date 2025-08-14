package main

import "fmt"

func main() {

	// if语句
	var a int = 3
	if a == 3 {
		fmt.Println("a等于3")
	}

	if b := 3; b == 3 {
		fmt.Println("b等于3")
	}

	// if else 语句
	if a == 3 {
		fmt.Println("a等于3")
	} else {
		fmt.Println("a不等于3")
	}

	// if else if 嵌套语句
	if a == 3 {
		fmt.Println("a等于3")
	} else if a == 4 {
		fmt.Println("a等于4")
	} else if a == 5 {
		fmt.Println("a等于5")
	} else {
		fmt.Println("a不等于3也不等于4也不等于5")
	}
}
