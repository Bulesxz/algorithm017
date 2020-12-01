package main

/*
子问题
 ......() 前面i-2 子问题
 ......))  子问题 .......((.....))
*/
func longestValidParentheses(s string) int {
	if len(s) == 0 {
		return 0
	}
	dp := make([]int, len(s))
	dp[0] = 0
	max := dp[0]
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			// ......() 前面i-2 子问题
			if s[i-1] == '(' {
				pre := 0
				if i-2 >= 0 {
					pre = dp[i-2]
				}
				dp[i] = pre + 2
			} else {
				/// ......))
				left := i - dp[i-1] - 1
				// fmt.Println(left-1,dp[left-1])
				if left >= 0 && s[left] == '(' {
					pre := 0
					if left-1 >= 0 {
						pre = dp[left-1]
					}
					dp[i] = (dp[i-1] + 2) + pre
					// fmt.Println(i,dp[i])
				}
			}
		}
		if dp[i] > max {
			max = dp[i]
		}
		// fmt.Println(dp)
	}
	return max
}
