package _79_numSquares

import "math"

// 279. 完全平方数
// 给你一个整数 n ，返回 和为 n 的完全平方数的最少数量 。
// 完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。
// 例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。
//
// 示例 1：
// 输入：n = 12
// 输出：3
// 解释：12 = 4 + 4 + 4
//
// 示例 2：
// 输入：n = 13
// 输出：2
// 解释：13 = 4 + 9
//
// 提示：
// 1 <= n <= 104

// 定义dp[i] 代表i的完全平方数最少的数量
// 则有 dp[i] = min(1 + dp[i-m*m])  (对所有 m ∈ [1, sqrt(i)] 求最小值)
// dp[0] = 0, dp[1] = 1
// 所求即为 dp[n]
func numSquares(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 1
	for i := 2; i <= n; i++ {
		mi := int(math.Sqrt(float64(i)))
		minComb := math.MaxInt
		for j := 1; j <= mi; j++ {
			minComb = min(minComb, 1+dp[i-j*j])
		}
		dp[i] = minComb
	}
	return dp[n]
}
