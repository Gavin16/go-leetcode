package _31_minFallingPathSum

import "math"

// 931. 下降路径最小和
// 中等
// 相关标签
// premium lock icon
// 相关企业
// 给你一个 n x n 的 方形 整数数组 matrix ，请你找出并返回通过 matrix 的下降路径 的 最小和 。
//
// 下降路径 可以从第一行中的任何元素开始，并从每一行中选择一个元素。在下一行选择的元素和当前行所选元素最多相隔一列（即位于正下方或者沿对角线向左或者向右的第一个元素）。
// 具体来说，位置 (row, col) 的下一个元素应当是 (row + 1, col - 1)、(row + 1, col) 或者 (row + 1, col + 1) 。
//
// 示例 1：
//
// 输入：matrix = [[2,1,3],[6,5,4],[7,8,9]]
// 输出：13
// 解释：如图所示，为和最小的两条下降路径
// 示例 2：
//
// 输入：matrix = [[-19,57],[-40,-5]]
// 输出：-59
// 解释：如图所示，为和最小的下降路径
//
// 提示：
//
// n == matrix.length == matrix[i].length
// 1 <= n <= 100
// -100 <= matrix[i][j] <= 100

// f[i][j] = min(f[i-1][j-1], f[i-1][j], f[i-1][j+1]) + matrix[i][j]
func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	if n == 1 {
		return matrix[0][0]
	}
	f := make([][]int, n)
	for idx := range f {
		f[idx] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		f[0][i] = matrix[0][i]
	}
	for i := 1; i < n; i++ {
		for j := 0; j < n; j++ {
			if j > 0 && j < n-1 {
				f[i][j] = min(f[i-1][j-1], f[i-1][j],
					f[i-1][j+1]) + matrix[i][j]
			} else if j == 0 {
				f[i][j] = min(f[i-1][j], f[i-1][j+1]) + matrix[i][j]
			} else if j == n-1 {
				f[i][j] = min(f[i-1][j], f[i-1][j-1]) + matrix[i][j]
			}
		}
	}
	minSum := math.MaxInt
	for j := 1; j < n; j++ {
		minSum = min(minSum, f[n-1][j], f[n-1][j-1])
	}
	return minSum
}
