package main

/*
1021. 删除最外层的括号
https://leetcode-cn.com/problems/remove-outermost-parentheses/
*/
func removeOuterParentheses(S string) string {
	left := 0
	right := 0
	result := []byte{}
	for i := 0; i < len(S); i++ {
		if S[i] == '(' {
			left++
		} else {
			right++
		}

		if left == 1 && right == 0 {
			//第一个左括号
			//not do
		} else if left == right {
			left = 0
			right = 0
		} else {
			result = append(result, S[i])
		}
	}
	return string(result)
}
