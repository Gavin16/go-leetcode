package _20_minimumTotal

import "math"

// 120. 三角形最小路径和
func minimumTotal(triangle [][]int) int {
	// 采用滚动数组，从后向前覆盖方式 可实现O(n) 空间复杂度
	N := len(triangle)
	dp := make([]int, N)
	dp[0] = triangle[0][0]
	for i := 1; i < N; i++ {
		for j := i; j >= 0; j-- {
			if j == 0 {
				dp[j] += triangle[i][j]
			} else if j == i {
				dp[j] = dp[j-1] + triangle[i][j]
			} else {
				dp[j] = min(dp[j-1], dp[j]) + triangle[i][j]
			}
		}
	}
	minSum := math.MaxInt
	for _, num := range dp {
		minSum = min(minSum, num)
	}
	return minSum
}
