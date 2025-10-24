package backtrace

import "math"

// 264. 丑数 II
// 给你一个整数 n ，请你找出并返回第 n 个 丑数 。
//
// 丑数 就是质因子只包含 2、3 和 5 的正整数。
//
// 示例 1：
//
// 输入：n = 10
// 输出：12
// 解释：[1, 2, 3, 4, 5, 6, 8, 9, 10, 12] 是由前 10 个丑数组成的序列。
// 示例 2：
//
// 输入：n = 1
// 输出：1
// 解释：1 通常被视为丑数。
//
// 提示：
//
// 1 <= n <= 1690

func nthUglyNumber(n int) int {
	idx := make([]int, 3)
	primes := []int{2, 3, 5}
	dp := make([]int, n+1)
	dp[1] = 1
	idx[0], idx[1], idx[2] = 1, 1, 1
	for i := 2; i <= n; i++ {
		minUgly := math.MaxInt
		for j := 0; j < len(primes); j++ {
			minUgly = min(minUgly, primes[j]*dp[idx[j]])
		}
		dp[i] = minUgly
		for k := 0; k < len(primes); k++ {
			if primes[k]*dp[idx[k]] == minUgly {
				idx[k]++
			}
		}
	}
	return dp[n]
}
