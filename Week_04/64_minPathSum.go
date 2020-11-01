package main

/*
64. 最小路径和
https://leetcode-cn.com/problems/minimum-path-sum/
*/
func dpMinPathSum(m, n int, grid [][]int) int {
	//1. 递归终止
	if m == 0 && n == 0 {
		return grid[m][n]
	}
	//2. 子问题

	if m == 0 {
		return dpMinPathSum(m, n-1, grid) + grid[m][n]
	}
	if n == 0 {
		return dpMinPathSum(m-1, n, grid) + grid[m][n]
	}
	a := dpMinPathSum(m, n-1, grid)
	b := dpMinPathSum(m-1, n, grid)
	if a < b {
		return a + grid[m][n]
	} else {
		return b + grid[m][n]
	}

}
func minPathSum(grid [][]int) int {
	m := len(grid) - 1
	n := len(grid[m]) - 1
	return dpMinPathSum(m, n, grid)
}
