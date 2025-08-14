package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// 1. 往文件中写入内容
func WriteDataToFile() {
	file, err := os.Create("./test.txt")
	if err != nil {
		fmt.Println(err)
	}
	// 在main函数结束时关闭文件流
	defer file.Close()

	for i := 0; i < 10; i++ {
		outstr := fmt.Sprintf("%s:%d", "hello go", i)
		file.WriteString(outstr)
		file.Write([]byte("\n"))
	}
	file.Write([]byte("write end！"))
}

// 2. 从文件中读取内容
func ReadDataFromFile() {
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	buf := make([]byte, 1024)
	var content strings.Builder

	for {
		n, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				break // 正常结束
			}
			fmt.Println("读取错误:", err)
			return
		} else {
			// 如果n == 0 说明已经读取完了
			if n == 0 {
				break
			}
		}
		content.Write(buf[:n])
	}
	// 打印数组结果
	fmt.Println("文件完整内容:")
	fmt.Println(content.String())
}

// 文件操作
func main() {
	//WriteDataToFile()

	ReadDataFromFile()
}
