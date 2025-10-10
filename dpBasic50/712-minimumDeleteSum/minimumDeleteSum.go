package _12_minimumDeleteSum

// 712. 两个字符串的最小ASCII删除和
// 给定两个字符串s1 和 s2，返回 使两个字符串相等所需删除字符的 ASCII 值的最小和 。
//
// 示例 1:
//
// 输入: s1 = "sea", s2 = "eat"
// 输出: 231
// 解释: 在 "sea" 中删除 "s" 并将 "s" 的值(115)加入总和。
// 在 "eat" 中删除 "t" 并将 116 加入总和。
// 结束时，两个字符串相等，115 + 116 = 231 就是符合条件的最小和。
// 示例 2:
//
// 输入: s1 = "delete", s2 = "leet"
// 输出: 403
// 解释: 在 "delete" 中删除 "dee" 字符串变成 "let"，
// 将 100[d]+101[e]+101[e] 加入总和。在 "leet" 中删除 "e" 将 101[e] 加入总和。
// 结束时，两个字符串都等于 "let"，结果即为 100+101+101+101 = 403 。
// 如果改为将两个字符串转换为 "lee" 或 "eet"，我们会得到 433 或 417 的结果，比答案更大。
// 提示:
//
// 0 <= s1.length, s2.length <= 1000
// s1 和 s2 由小写英文字母组成
func minimumDeleteSum(s1 string, s2 string) int {
	sum1, sum2 := 0, 0
	for i := 0; i < len(s1); i++ {
		sum1 += int(s1[i])
	}
	for j := 0; j < len(s2); j++ {
		sum2 += int(s2[j])
	}
	m, n := len(s1), len(s2)
	dp := make([][]int, m+1)
	for idx := range dp {
		dp[idx] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if s1[i] == s2[j] {
				dp[i+1][j+1] = max(dp[i][j]+int(s1[i]), dp[i][j+1], dp[i+1][j])
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	return sum1 + sum2 - 2*dp[m][n]
}
