package _22_coinChange

import "math"

// 322. 零钱兑换
// 给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
//
// 计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
//
// 你可以认为每种硬币的数量是无限的。
//
// 示例 1：
//
// 输入：coins = [1, 2, 5], amount = 11
// 输出：3
// 解释：11 = 5 + 5 + 1
// 示例 2：
//
// 输入：coins = [2], amount = 3
// 输出：-1
// 示例 3：
//
// 输入：coins = [1], amount = 0
// 输出：0
//
// 提示：
//
// 1 <= coins.length <= 12
// 1 <= coins[i] <= 231 - 1
// 0 <= amount <= 104

// 1. 难度: 在范围内, 2.目标: 明确, 3.专注度: 专注, 4.反馈: 待参考
// 深度优先搜索 -- 加入memo 避免超时
func coinChange(coins []int, amount int) int {
	memo := make([]int, amount+1)
	return dfs(coins, amount, memo)
}

func dfs(coins []int, amount int, memo []int) int {
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	// 对于所有计算过的情况，都需要返回
	if memo[amount] != 0 {
		return memo[amount]
	}
	temp := math.MaxInt
	for j := 0; j < len(coins); j++ {
		if amount < coins[j] {
			continue
		}
		deep := dfs(coins, amount-coins[j], memo)
		if deep >= 0 {
			temp = min(temp, deep+1)
		}
	}
	if temp == math.MaxInt {
		memo[amount] = -1
	} else {
		memo[amount] = temp
	}
	return memo[amount]
}

// 动态规划解法
// 这里金额取值范围最小，希望针对给定的 coins 列表，找出满足拼凑amount 金额 能找出的最小的硬币数
// 金额 i < 0, 硬币数为0
// 金额 i = 0, 硬币数为0
// 金额 i > 0, dp[i] = min(dp[i - val_j ] + 1) 其中val_j 为coins 中第j个硬币 j ∈ [0, len(coins)]
func coinChange1(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	dict := make(map[int]int)
	for _, size := range coins {
		dict[size] = 1
	}

	dp := make([]int, amount+1)
	maxFill := 0
	for _, val := range coins {
		if val > amount {
			continue
		}
		if dict[val] == 1 {
			dp[val] = 1
			maxFill = max(maxFill, val)
		} else {
			dp[val] = -1
		}
	}
	for k := maxFill + 1; k <= amount; k++ {
		currMin := amount + 1
		for j := 0; j < len(coins); j++ {
			if coins[j] > k {
				continue
			}
			if dp[k-coins[j]] >= 1 {
				currMin = min(currMin, dp[k-coins[j]]+1)
			}
		}
		if currMin > amount {
			dp[k] = -1
		} else {
			dp[k] = currMin
		}
	}
	return dp[amount]
}

func coinChange2(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		minCount := math.MaxInt
		for j := 0; j < len(coins); j++ {
			coinSize := coins[j]
			if coinSize <= i && dp[i-coinSize] < minCount {
				minCount = dp[i-coinSize] + 1
			}
		}
		dp[i] = minCount
	}
	if dp[amount] == math.MaxInt {
		return -1
	} else {
		return dp[amount]
	}
}
