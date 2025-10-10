package _16_longestPalindromeSubseq

// 516. 最长回文子序列
// 给你一个字符串 s ，找出其中最长的回文子序列，并返回该序列的长度。
//
// 子序列定义为：不改变剩余字符顺序的情况下，删除某些字符或者不删除任何字符形成的一个序列。
//
// 示例 1：
//
// 输入：s = "bbbab"
// 输出：4
// 解释：一个可能的最长回文子序列为 "bbbb" 。
// 示例 2：
//
// 输入：s = "cbbd"
// 输出：2
// 解释：一个可能的最长回文子序列为 "bb" 。
//
// 提示：
//
// 1 <= s.length <= 1000
// s 仅由小写英文字母组成

func longestPalindromeSubseq(s string) int {
	N := len(s)
	dp := make([][]int, N)
	for idx := range dp {
		dp[idx] = make([]int, N)
		dp[idx][idx] = 1
	}
	for k := 1; k < N; k++ {
		for i, j := 0, k; i < N-k && j <= N-1; {
			if j-i == 1 {
				if s[i] == s[j] {
					dp[i][j] = 2
				} else {
					dp[i][j] = 1
				}
			} else {
				temp := dp[i+1][j-1]
				if s[i] == s[j] {
					temp += 2
				}
				dp[i][j] = max(temp, dp[i+1][j], dp[i][j-1])
			}
			i, j = i+1, j+1
		}
	}
	return dp[0][N-1]
}
