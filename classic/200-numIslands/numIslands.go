package _00_numIslands

// 200. 岛屿数量
// 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
//
// 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
//
// 此外，你可以假设该网格的四条边均被水包围。
//
// 示例 1：
//
// 输入：grid = [
//
//	["1","1","1","1","0"],
//	["1","1","0","1","0"],
//	["1","1","0","0","0"],
//	["0","0","0","0","0"]
//
// ]
// 输出：1
// 示例 2：
//
// 输入：grid = [
//
//	["1","1","0","0","0"],
//	["1","1","0","0","0"],
//	["0","0","1","0","0"],
//	["0","0","0","1","1"]
//
// ]
// 输出：3
//
// 提示：
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 300
// grid[i][j] 的值为 '0' 或 '1'

// 击败: 30.13%
func numIslands(grid [][]byte) int {
	rows, cols := len(grid), len(grid[0])

	var diffuse func(board [][]byte, x, y int)
	diffuse = func(board [][]byte, x, y int) {
		if x < 0 || x >= rows || y < 0 || y >= cols {
			return
		}
		if board[x][y] == '0' || board[x][y] == '2' {
			return
		}
		board[x][y] = '2'
		diffuse(grid, x-1, y)
		diffuse(grid, x+1, y)
		diffuse(grid, x, y-1)
		diffuse(grid, x, y+1)
	}

	count := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				diffuse(grid, i, j)
				count += 1
			}
		}
	}
	return count
}
