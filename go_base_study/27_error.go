package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divide by zero")
	}
	return a / b, nil
}

func main() {

	var err error = errors.New("这是一个错误")
	fmt.Println(err)

	var err1 error = fmt.Errorf("这是一个错误:%s", "a normal err1")
	fmt.Println(err1)

	// 调用函数
	res, err := divide(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

	// 调用函数
	result, err1 := divide(10, 5)
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(result)
	}
}
