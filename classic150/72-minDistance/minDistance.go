package _2_minDistance

// 72. 编辑距离
// 给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数  。
//
// 你可以对一个单词进行如下三种操作：
//
// 插入一个字符
// 删除一个字符
// 替换一个字符
//
// 示例 1：
//
// 输入：word1 = "horse", word2 = "ros"
// 输出：3
// 解释：
// horse -> rorse (将 'h' 替换为 'r')
// rorse -> rose (删除 'r')
// rose -> ros (删除 'e')
// 示例 2：
//
// 输入：word1 = "intention", word2 = "execution"
// 输出：5
// 解释：
// intention -> inention (删除 't')
// inention -> enention (将 'i' 替换为 'e')
// enention -> exention (将 'n' 替换为 'x')
// exention -> exection (将 'n' 替换为 'c')
// exection -> execution (插入 'u')
//
// 提示：
//
// 0 <= word1.length, word2.length <= 500
// word1 和 word2 由小写英文字母组成

// 定义 f(i,j) 为从word1 的第i个位置到第word2 的第j个位置变换所需要的 最少操作数
// 若word1[i] = word2[j] 则 f(i, j) = f(i-1, j-1)   (i, j)
// 若word1[i] != word2[j] 则 f(i, j) = min(f(i-1, j-1) + 1, f(i-1, j) + 1)
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	if m*n == 0 {
		return n + m
	}

	dp := make([][]int, m+1)
	for idx := range dp {
		dp[idx] = make([]int, n+1)
	}
	dp[0][0] = 0
	for k := 1; k <= m; k++ {
		dp[k][0] = k
	}
	for k := 1; k <= n; k++ {
		dp[0][k] = k
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if word2[i-1] == word1[j-1] {
				dp[j][i] = min(dp[j-1][i-1], dp[j-1][i]+1, dp[j][i-1]+1)
			} else {
				dp[j][i] = min(dp[j-1][i-1]+1, dp[j-1][i]+1, dp[j][i-1]+1)
			}
		}
	}
	return dp[m][n]
}
