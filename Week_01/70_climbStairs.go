package main

/*
70 爬楼梯
https://leetcode-cn.com/problems/climbing-stairs/
*/

//递归版本 2**n
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	return climbStairs(n-1) + climbStairs(n-2)
}

//迭代版本
func climbStairs2(n int) int {
	if n <= 2 {
		return n
	}
	f1 := 1
	f2 := 2
	f3 := -1
	for i := 3; i <= n; i++ {
		f3 = f1 + f2
		f1 = f2
		f2 = f3
	}
	return f3
}
