package main

import (
	"encoding/json"
	"fmt"
)

// json.Unmarshal()  json转为对象

// 定义一个结构体
type itTest struct {
	Subjects []string
	IsOk     bool
	IsTrue   bool
}

type itTest2 struct {
	Subjects []string `json:"subjects"`
}

func main() {

	b1 := []byte(`{
		"subjects": [
		"go",
		"java",
		"python"
		],
		"IsOk": true,
		"isTrue": true
	}`)

	var itTest itTest

	// 记住这个& 必须先定义才可以获取
	// TODO 千万注意 字段名要是大写 否则的话不会被正确的解析

	err := json.Unmarshal(b1, &itTest)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(itTest)
	}

	// 只想要一个属性subjects
	var itTest2 itTest2
	err = json.Unmarshal(b1, &itTest2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(itTest2)
	}

	// 解析到interface 他就是一个map
	var itTest3 interface{}
	err = json.Unmarshal(b1, &itTest3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(itTest3)
	}
	// 从interface中获取数据
	m := itTest3.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, vv)
		case bool:
			fmt.Println(k, vv)
		case float64:
			fmt.Println(k, vv)
		case []interface{}:
			fmt.Println(k, vv)
			for _, v := range vv {
				fmt.Println(v)
			}
		}
	}
}
