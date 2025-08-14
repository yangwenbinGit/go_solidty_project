package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	o := make(chan bool)

	go func() {
		select {
		case v := <-ch:
			fmt.Println(v)
		case <-time.After(5 * time.Second):
			fmt.Println("timeout")
			o <- true
			break
		}
	}()

	value := <-o
	fmt.Println(value)
}
