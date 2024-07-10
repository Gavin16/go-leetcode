package _89_gameOfLife

// 289. 生命游戏
// 根据 百度百科 ， 生命游戏 ，简称为 生命 ，是英国数学家约翰·何顿·康威在 1970 年发明的细胞自动机。
//
// 给定一个包含 m × n 个格子的面板，每一个格子都可以看成是一个细胞。每个细胞都具有一个初始状态：
// 1 即为 活细胞 （live），或 0 即为 死细胞 （dead）。每个细胞与其八个相邻位置（水平，垂直，对角线）的细胞都遵循以下四条生存定律：
//
// 如果活细胞周围八个位置的活细胞数少于两个，则该位置活细胞死亡；
// 如果活细胞周围八个位置有两个或三个活细胞，则该位置活细胞仍然存活；
// 如果活细胞周围八个位置有超过三个活细胞，则该位置活细胞死亡；
// 如果死细胞周围正好有三个活细胞，则该位置死细胞复活；
//
// 下一个状态是通过将上述规则同时应用于当前状态下的每个细胞所形成的，其中细胞的出生和死亡是同时发生的。
// 给你 m x n 网格面板 board 的当前状态，返回下一个状态。
//
// 示例 1：
//
// 输入：board = [[0,1,0],[0,0,1],[1,1,1],[0,0,0]]
// 输出：[[0,0,0],[1,0,1],[0,1,1],[0,1,0]]
// 示例 2：
//
// 输入：board = [[1,1],[1,0]]
// 输出：[[1,1],[1,1]]
//
// 提示：
//
// m == board.length
// n == board[i].length
// 1 <= m, n <= 25
// board[i][j] 为 0 或 1
//
// 进阶：
//
// 你可以使用原地算法解决本题吗？请注意，面板上所有格子需要同时被更新：你不能先更新某些格子，然后使用它们的更新后的值再更新其他格子。
// 本题中，我们使用二维数组来表示面板。原则上，面板是无限的，但当活细胞侵占了面板边界时会造成问题。你将如何解决这些问题？

// 原地解决问题可以考虑增加 状态
// 对于 取值为 0, 1 的分别代表 死细胞 和 活细胞
// 对于 取值不为 0,1 的 通过变化编码来记录
// 具体表现为:  死 -> 活 (0 -> 2)  活 -> 死(1 -> 3) 死 -> 死 (0 -> 0) 活 -> 活(1 -> 1)
// 当进行遍历判断时，通过  board[ci][cj]%2 == 1 可以获得它之前的状态，在计算得到的就都是之前的状态
// 最后扫描一遍board 将状态统一

// 击败: 16.80%
func gameOfLife(board [][]int) {
	rows, cols := len(board), len(board[0])
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1},
		{0, 1}, {-1, -1}, {1, 1}, {-1, 1}, {1, -1}}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			ci, cj, liveCnt := i, j, 0
			for _, dir := range dirs {
				ci, cj = i+dir[0], j+dir[1]
				if ci < 0 || ci >= rows || cj < 0 || cj >= cols {
					continue
				}
				if board[ci][cj]%2 == 1 {
					liveCnt++
				}
			}
			if board[i][j] == 1 {
				switch liveCnt {
				case 0, 1, 4, 5, 6, 7, 8:
					board[i][j] = 3
				default:
					board[i][j] = board[i][j]
				}
			}
			if board[i][j] == 0 {
				if liveCnt == 3 {
					board[i][j] = 2
				}
			}
		}
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == 1 || board[i][j] == 2 {
				board[i][j] = 1
			} else {
				board[i][j] = 0
			}
		}
	}
}