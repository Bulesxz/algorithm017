package main

/*
22. 括号生成
https://leetcode-cn.com/problems/generate-parentheses/
*/

func backtraceGenerateParenthesis(right, left, n int, path string, res *[]string) {
	//1.满足条件
	if right == 0 && left == 0 {
		//左右括号都使用完了
		*res = append(*res, path)
		return
	}

	//2.选择列表
	if left > 0 { //还有左括号
		path = path + "(" //选择
		backtraceGenerateParenthesis(right, left-1, n, path, res)
		path = path[:len(path)-1] //撤销选择
	}

	if right > 0 && right > left {
		//剩余右括号大于左括号，即左括号使用的多余右括号
		path = path + ")" //选择
		backtraceGenerateParenthesis(right-1, left, n, path, res)
		path = path[:len(path)-1] //撤销选择
	}
}
func generateParenthesis(n int) []string {
	path := ""
	res := []string{}
	backtraceGenerateParenthesis(n, n, n, path, &res)
	return res
}
