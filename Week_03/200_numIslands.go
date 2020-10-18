package main

func dfs(i, j int, grid [][]byte) {

	//1.访问
	grid[i][j] = 0 //变成水

	rowSize := len(grid)
	colSize := len(grid[i])

	//2.遍历
	//1.右边
	if j+1 < colSize && grid[i][j+1] == '1' {
		dfs(i, j+1, grid)
	}
	//2.左边
	if j-1 >= 0 && grid[i][j-1] == '1' {
		dfs(i, j-1, grid)
	}

	//3.上
	if i-1 >= 0 && grid[i-1][j] == '1' {
		dfs(i-1, j, grid)
	}
	//4.下
	if i+1 < rowSize && grid[i+1][j] == '1' {
		dfs(i+1, j, grid)
	}

}
func numIslands(grid [][]byte) int {
	count := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				count++
				dfs(i, j, grid)
			}
		}

	}
	return count
}
