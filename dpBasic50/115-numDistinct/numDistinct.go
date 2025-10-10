package _15_numDistinct

// 115. 不同的子序列
// 给你两个字符串 s 和 t ，统计并返回在 s 的 子序列 中 t 出现的个数。
// 测试用例保证结果在 32 位有符号整数范围内。
//
// 示例 1：
//
// 输入：s = "rabbbit", t = "rabbit"
// 输出：3
// 解释：
// 如下所示, 有 3 种可以从 s 中得到 "rabbit" 的方案。
// rabb_it
// ra_bbit
// rab_bit
// 示例 2：
//
// 输入：s = "babgbag", t = "bag"
// 输出：5
// 解释：
// 如下所示, 有 5 种可以从 s 中得到 "bag" 的方案。
// babgbag
// babgbag
// babgbag
// babgbag
// babgbag
//
// 提示：
// 1 <= s.length, t.length <= 1000
// s 和 t 由英文字母组成

// dp[i][j] 代表 s[..,i]中所有子序列中包含 t[..,j] 的个数
// 若s[i] == t[j] 则 dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
// 若s[i] != t[j] 则 dp[i][j] = dp[i-1][j]
// dp[i][0] = 1,  对于 i < j dp[i][j] = 0
func numDistinct(s string, t string) int {
	m, n := len(s), len(t)
	if m < n {
		return 0
	}
	dp := make([][]int, m+1)
	for i, _ := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = 1
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= i && j <= n; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j] + dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[m][n]
}
