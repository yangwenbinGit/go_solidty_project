package main

import (
	"fmt"
)

func testA1() {
	fmt.Println("func TestA()")
}

func testB1() (err error) {
	defer func() { //在发生异常时，设置恢复
		if x := recover(); x != nil {

			//并用err变量接收错误信息，返回给调用者。x就是panic的参数
			err = fmt.Errorf("internal error: %v", x)
		}
	}()

	panic("func TestB(): panic")
}

func TestC1() {
	fmt.Println("func TestC()")
}

func main() {
	testA1()
	err := testB1()
	if err != nil {
		fmt.Println(err)
	}
	TestC1()
}
