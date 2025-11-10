package _29_longestIncreasingPath

import "slices"

// 329. 矩阵中的最长递增路径
// 给定一个 m x n 整数矩阵 matrix ，找出其中 最长递增路径 的长度。
//
// 对于每个单元格，你可以往上，下，左，右四个方向移动。 你 不能 在 对角线 方向上移动或移动到 边界外（即不允许环绕）。
//
// 示例 1：
//
// 输入：matrix = [[9,9,4],[6,6,8],[2,1,1]]
// 输出：4
// 解释：最长递增路径为 [1, 2, 6, 9]。
// 示例 2：
//
// 输入：matrix = [[3,4,5],[3,2,6],[2,2,1]]
// 输出：4
// 解释：最长递增路径是 [3, 4, 5, 6]。注意不允许在对角线方向上移动。
// 示例 3：
//
// 输入：matrix = [[1]]
// 输出：1
//
// 提示：
//
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 200
// 0 <= matrix[i][j] <= 231 - 1

func longestIncreasingPath(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	dirs := [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
	var dfs func(pi, pj int, i, j int)
	dfs = func(pi, pj int, i, j int) {
		if i >= m || i < 0 || j >= n || j < 0 {
			return
		}
		if matrix[i][j] <= matrix[pi][pj] {
			return
		}
		dp[i][j] = max(dp[i][j], dp[pi][pj]+1)
		for _, dir := range dirs {
			ni, nj := i+dir[0], j+dir[1]
			if ni >= m || ni < 0 || nj >= n || nj < 0 {
				continue
			}
			if dp[i][j] >= dp[ni][nj] {
				dfs(i, j, ni, nj)
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dp[i][j] == -1 {
				dp[i][j] = 1
				for _, dir := range dirs {
					ni, nj := i+dir[0], j+dir[1]
					dfs(i, j, ni, nj)
				}
			}
		}
	}
	maxLen := 0
	for _, row := range dp {
		maxLen = max(maxLen, slices.Max(row))
	}
	return maxLen
}
