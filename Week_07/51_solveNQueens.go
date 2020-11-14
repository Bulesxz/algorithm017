package main

/*
51. N 皇后
https://leetcode-cn.com/problems/n-queens/

回溯法:
剪枝
*/

func brackTraceSolveNQueens(row, n int, board [][]int, res *[][]string) {
	// fmt.Println(row, board)
	//1. 满足条件
	if row == n {
		r := []string{}
		for i := 0; i < n; i++ {
			s := ""
			for j := 0; j < n; j++ {
				if board[i][j] == 1 {
					s += "Q"
				} else {
					s += "."
				}
			}
			r = append(r, s)
		}

		*res = append(*res, r)
		return
	}
	//2. 列 选择列表
	for i := 0; i < n; i++ {
		if checkBorad(row, i, n, board) {
			board[row][i] = 1
			brackTraceSolveNQueens(row+1, n, board, res)
			board[row][i] = 0
		}
	}
}

//检查棋盘
func checkBorad(row, col, n int, board [][]int) bool {
	// for i := 0; i < col; i++ {
	// 	if board[row][col] == 1 { //行冲突
	// 		return false
	// 	}
	// } //不需要判断行，因为是一行一行往下走的
	for i := 0; i < row; i++ {
		if board[i][col] == 1 { //列冲突
			return false
		}
	}

	for i := 1; i <= row; i++ {
		if col-i >= 0 && board[row-i][col-i] == 1 { //左上冲突
			return false
		}
		if col+i < n && board[row-i][col+i] == 1 { //右上冲突
			return false
		}
	}
	return true
}
func solveNQueens(n int) [][]string {
	res := [][]string{}
	board := [][]int{} //初始化棋盘
	for i := 0; i < n; i++ {
		board = append(board, make([]int, n, n))
	}

	brackTraceSolveNQueens(0, n, board, &res)

	return res
}
