package _684_maxMoves

import "slices"

// 2684. 矩阵中移动的最大次数
// 给你一个下标从 0 开始、大小为 m x n 的矩阵 grid ，矩阵由若干 正 整数组成。
//
// 你可以从矩阵第一列中的 任一 单元格出发，按以下方式遍历 grid ：
//
// 从单元格 (row, col) 可以移动到 (row - 1, col + 1)、(row, col + 1) 和 (row + 1, col + 1) 三个单元格中任一满足值 严格 大于当前单元格的单元格。
// 返回你在矩阵中能够 移动 的 最大 次数。
//
// 示例 1：
//
// 输入：grid = [[2,4,3,5],[5,4,9,3],[3,4,2,11],[10,9,13,15]]
// 输出：3
// 解释：可以从单元格 (0, 0) 开始并且按下面的路径移动：
// - (0, 0) -> (0, 1).
// - (0, 1) -> (1, 2).
// - (1, 2) -> (2, 3).
// 可以证明这是能够移动的最大次数。
// 示例 2：
//
// 输入：grid = [[3,2,4],[2,1,9],[1,1,7]]
// 输出：0
// 解释：从第一列的任一单元格开始都无法移动。
//
// 提示：
//
// m == grid.length
// n == grid[i].length
// 2 <= m, n <= 1000
// 4 <= m * n <= 105
// 1 <= grid[i][j] <= 106

func maxMoves(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	for i := 0; i < m; i++ {
		dp[i][0] = 0
	}
	for j := 1; j < n; j++ {
		for i := 0; i < m; i++ {
			curr, left := grid[i][j], grid[i][j-1]
			fromLeft := calPath(dp[i][j-1], curr, left)
			if i > 0 && i < m-1 {
				fromLeftUp := calPath(dp[i-1][j-1], curr, grid[i-1][j-1])
				fromLeftDown := calPath(dp[i+1][j-1], curr, grid[i+1][j-1])
				dp[i][j] = max(dp[i][j], fromLeft, fromLeftUp, fromLeftDown)
			} else if i == 0 {
				fromLeftDown := calPath(dp[i+1][j-1], curr, grid[i+1][j-1])
				dp[i][j] = max(dp[i][j], fromLeft, fromLeftDown)
			} else if i == m-1 {
				fromLeftUp := calPath(dp[i-1][j-1], curr, grid[i-1][j-1])
				dp[i][j] = max(dp[i][j], fromLeft, fromLeftUp)
			}
		}
	}
	maxMov := 0
	for _, row := range dp {
		maxMov = max(maxMov, slices.Max(row))
	}
	return maxMov
}

func calPath(prev int, a, b int) int {
	if a > b && prev != -1 {
		return prev + 1
	}
	return -1
}

// 使用广度优先搜索
func maxMoves1(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	//queue := make([][]int, 0)
	q1 := make(map[int]bool)
	for i := 0; i < m; i++ {
		q1[i] = true
	}
	for j := 1; j < n; j++ {
		q2 := make(map[int]bool)
		for i, _ := range q1 {
			for k := i - 1; k <= i+1; k++ {
				if k < 0 || k >= m {
					continue
				}
				if grid[k][j] > grid[i][j-1] {
					q2[k] = true
				}
			}
		}
		q1 = q2
		if len(q1) == 0 {
			return j - 1
		}
	}
	return n - 1
}
