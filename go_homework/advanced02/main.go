package main

import "fmt"

/**
题目： 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：
左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。


示例 1：
输入：s = "()"
输出：true

示例 2：
输入：s = "()[]{}"
输出：true

示例 3：
输入：s = "(]"
输出：false

示例 4：
输入：s = "([])"
输出：true

s 仅由括号 '()[]{}' 组成
*/

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	// 定义一个栈
	stack := make([]byte, 0)
	// 遍历字符串
	for i := 0; i < len(s); i++ {
		// 如果是左括号 就入栈
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else {
			// 如果是右括号 就出栈
			if len(stack) == 0 {
				return false
			}
			// 出栈
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// 判断是否匹配
			if (s[i] == ')' && top != '(') || (s[i] == ']' && top != '[') || (s[i] == '}' && top != '{') {
				return false
			}
		}
	}
	// 判断栈是否为空
	if len(stack) > 0 {
		return false
	}
	return true
}

func main() {
	//s := "()[]{}"
	s := "(}"
	fmt.Println(isValid(s))
}
