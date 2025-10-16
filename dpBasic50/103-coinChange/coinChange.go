package _03_coinChange

import "math"

func coinChange(coins []int, amount int) int {
	// dp[i] 代表金额 i 所需的最少硬币数
	// dp[i] = min(dp[i-c]+1, dp[i])
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		minNum := math.MaxInt
		for _, c := range coins {
			if i < c || dp[i-c] == -1 {
				continue
			}
			minNum = min(minNum, dp[i-c]+1)
		}
		if minNum == math.MaxInt {
			dp[i] = -1
		} else {
			dp[i] = minNum
		}
	}
	return dp[amount]
}
