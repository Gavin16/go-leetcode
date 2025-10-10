package _31_minFallingPathSum

import "math"

// 931. 下降路径最小和
func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	dp := make([][]int, n)
	for idx := range dp {
		dp[idx] = make([]int, n)
		dp[0][idx] = matrix[0][idx]
	}
	for i := 1; i < n; i++ {
		for j := 0; j < n; j++ {
			if j == 0 {
				dp[i][j] = min(dp[i-1][j], dp[i-1][j+1]) + matrix[i][j]
			} else if j == n-1 {
				dp[i][j] = min(dp[i-1][j], dp[i-1][j-1]) + matrix[i][j]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i-1][j-1], dp[i-1][j+1]) + matrix[i][j]
			}
		}
	}
	minSum := math.MaxInt
	for i := 0; i < n; i++ {
		minSum = min(minSum, dp[n-1][i])
	}
	return minSum
}
