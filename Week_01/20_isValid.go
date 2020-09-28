package main

/*
20 有效括号
https://leetcode-cn.com/problems/valid-parentheses/description/
*/
func isValid(s string) bool {
	stack := make([]string, len(s))
	stackTop := -1
	for _, c := range s {
		if string(c) == "[" || string(c) == "(" || string(c) == "{" || stackTop == -1 {
			stackTop++
			stack[stackTop] = string(c)
		} else if string(c) == "}" {
			if stack[stackTop] == "{" {
				stackTop--
			} else {
				return false
			}
		} else if string(c) == ")" {
			if stack[stackTop] == "(" {
				stackTop--
			} else {
				return false
			}
		} else if string(c) == "]" {
			if stack[stackTop] == "[" {
				stackTop--
			} else {
				return false
			}
		} else {
			return false
		}
	}
	if stackTop == -1 {
		return true
	}
	return false
}
