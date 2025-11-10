package LCR100_minimumTotal

import "slices"

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	dp := make([]int, n)
	dp[0] = triangle[0][0]
	for i := 1; i < n; i++ {
		for j := len(triangle[i]) - 1; j >= 0; j-- {
			v := triangle[i][j]
			if j == 0 {
				dp[j] = dp[j] + v
			} else if j == i {
				dp[j] = dp[j-1] + v
			} else {
				dp[j] = min(dp[j-1], dp[j]) + v
			}
		}
	}
	return slices.Min(dp)
}
