package backtrace

import (
	"math"
)

// 313. 超级丑数
// 超级丑数 是一个正整数，并满足其所有质因数都出现在质数数组 primes 中。
// 给你一个整数 n 和一个整数数组 primes ，返回第 n 个 超级丑数 。
// 题目数据保证第 n 个 超级丑数 在 32-bit 带符号整数范围内。
//
// 示例 1：
// 输入：n = 12, primes = [2,7,13,19]
// 输出：32
// 解释：给定长度为 4 的质数数组 primes = [2,7,13,19]，前 12 个超级丑数序列为：[1,2,4,7,8,13,14,16,19,26,28,32] 。
//
// 示例 2：
// 输入：n = 1, primes = [2,3,5]
// 输出：1
// 解释：1 不含质因数，因此它的所有质因数都在质数数组 primes = [2,3,5] 中。
//
// 提示：
// 1 <= n <= 105
// 1 <= primes.length <= 100
// 2 <= primes[i] <= 1000
// 题目数据 保证 primes[i] 是一个质数
// primes 中的所有值都 互不相同 ，且按 递增顺序 排列

// 定义dp[i] 代表第i个丑数的值
// 我们知道 dp[i] 的取值一定等于 dp 数组中i 之前的某个值乘以 primes 中某个j 得到
//   dp[i] 应该是dp[i] 之前范围内 挑选出来几个值分别于 primes[j] 相乘, 得到的最小的值作为dp[i]
// 使用idx[j] 记录每个与primes 中第j个质数相乘丑数在dp中的下标。
// 每次计算时找出 dp[idx[j]] * primes[j] j ∈ [0, len(primes))
// 每次计算完之后 对应j位置的 idx[j] 因为已经用过了，dp[i+1], dp[i+1] 再次用来乘以 primes[j] 的那个数应该是dp中的下一个丑数了

func nthSuperUglyNumber(n int, primes []int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	idx := make([]int, len(primes))
	for i := range idx {
		idx[i] = 1
	}
	for i := 2; i <= n; i++ {
		minMul := math.MaxInt
		for j := 0; j < len(primes); j++ {
			curr := dp[idx[j]] * primes[j]
			if curr > dp[i-1] {
				minMul = min(minMul, dp[idx[j]]*primes[j])
			}
		}
		dp[i] = minMul
		for k := 0; k < len(primes); k++ {
			if dp[idx[k]]*primes[k] == minMul {
				idx[k] += 1
			}
		}
	}
	return dp[n]
	//dp := make([]int, n+1)
	//dp[1] = 1
	//idx := make([]int, len(primes))
	//for i := range idx {
	//	idx[i] = 1
	//}
	//for i := 2; i <= n; i++ {
	//	minMul := math.MaxInt
	//	minJ := 0
	//	for j := len(primes) - 1; j >= 0; j-- {
	//		curr := dp[idx[j]] * primes[j]
	//		if curr <= minMul && curr > dp[i-1] {
	//			minMul = dp[idx[j]] * primes[j]
	//			minJ = j
	//		}
	//	}
	//	dp[i] = minMul
	//	idx[minJ] += 1
	//}
	//return dp[n]
}
