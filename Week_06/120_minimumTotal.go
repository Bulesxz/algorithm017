package main

/*
120. 三角形最小路径和
https://leetcode-cn.com/problems/triangle/description/

f(i,j)  = min ( f(i+1,j), f(i+1,j+1) )
*/
func minimumTotal(triangle [][]int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		} else {
			return b
		}
	}
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := i; j >= 0; j-- {
			triangle[i][j] = min(triangle[i+1][j], triangle[i+1][j+1]) + triangle[i][j]
		}
	}
	return triangle[0][0]
}
