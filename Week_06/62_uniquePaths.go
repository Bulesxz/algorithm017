package main

/*
62. 不同路径
https://leetcode-cn.com/problems/unique-paths/
*/

/*
1. 重复性子问题
	1. 下一个格子 只能向右 (列-1) 或者向下(行-1)
2. 状态定义
	1. i,j为格子的下标
	2. 起始点 为左上角(m,n), 终点为右下角(0,0)
	3. f(i,j) = f(i-1,j) + f(i,j-1)
3. 递推方程
	dp[i,j] = dp[i-1,j] + dp[i,j-1]
	边界:
	dp[i][0] = 1
	dp[0][j] = 1
*/
func uniquePaths(m int, n int) int {
	var dp = make([][]int, m, m)
	for i := 0; i < m; i++ {
		carray := make([]int, n, n)
		dp[i] = carray
	}
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

// 我们看上面二维数组的递推公式，当前坐标的值只和左边与上面的值有关，和其他的无关，这样二维数组造成大量的空间浪费，所以我们可以把它改为一维数组。
func uniquePaths2(m int, n int) int {

	dp := make([]int, n, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] = dp[j] + dp[j-1]
		}
	}
	return dp[n-1]
}
