package _30_solve

// 130. 被围绕的区域
// 给你一个 m x n 的矩阵 board ，由若干字符 'X' 和 'O' 组成，捕获 所有 被围绕的区域：
//
// 连接：一个单元格与水平或垂直方向上相邻的单元格连接。
// 区域：连接所有 'O' 的单元格来形成一个区域。
// 围绕：如果您可以用 'X' 单元格 连接这个区域，并且区域中没有任何单元格位于 board 边缘，则该区域被 'X' 单元格围绕。
// 通过将输入矩阵 board 中的所有 'O' 替换为 'X' 来 捕获被围绕的区域。
//
// 示例 1：
//
// 输入：board = [["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]
// 输出：[["X","X","X","X"],["X","X","X","X"],["X","X","X","X"],["X","O","X","X"]]
//
// 解释：
// 在上图中，底部的区域没有被捕获，因为它在 board 的边缘并且不能被围绕。
//
// 示例 2：
// 输入：board = [["X"]]
// 输出：[["X"]]
//
// 提示：
//
// m == board.length
// n == board[i].length
// 1 <= m, n <= 200
// board[i][j] 为 'X' 或 'O'

// 扫描边缘找出为 'O'的元素, 并做深度优先搜索替换 'O' 为 'Y'
// 扫描完之后将所有的 'O' 替换为 'X' 并将 'Y'替换回 'O'
// 击败: 78.65%
func solve(board [][]byte) {
	if len(board) < 3 || len(board[0]) < 3 {
		return
	}
	var dfsReplace func(grids [][]byte, i, j int)
	dfsReplace = func(grids [][]byte, i, j int) {
		if i < 0 || j < 0 || i >= len(grids) || j >= len(grids[0]) {
			return
		}
		if grids[i][j] != 'O' {
			return
		}
		grids[i][j] = 'Y'
		dfsReplace(grids, i, j+1)
		dfsReplace(grids, i, j-1)
		dfsReplace(grids, i+1, j)
		dfsReplace(grids, i-1, j)
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if i == 0 || i == len(board)-1 ||
				j == 0 || j == len(board[0])-1 {
				if board[i][j] == 'O' {
					dfsReplace(board, i, j)
				}
			}
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'Y' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}
