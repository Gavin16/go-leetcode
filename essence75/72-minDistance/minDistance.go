package _2_minDistance

// 72. 编辑距离
// 给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数  。
//
// 你可以对一个单词进行如下三种操作：
// 插入一个字符
// 删除一个字符
// 替换一个字符
//
// 示例 1：
// 输入：word1 = "horse", word2 = "ros"
// 输出：3
// 解释：
// horse -> rorse (将 'h' 替换为 'r')
// rorse -> rose (删除 'r')
// rose -> ros (删除 'e')
//
// 示例 2：
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

// 定义dp(i, j) 为从word1 的第i个字符 到 word2 的第j个字符位置 转换所需要的操作数
// 则 dp(i,j) = dp(i-1, j-1) + 1 若word1[i] != word2[j], 另外考虑到 dp(i,j) 也可能是从 dp(i-1, j) 或者 dp(i,j-1) 通过在word1 或者 word2 的后面
// 新增一个字符，如 dp(i,j) 可视为 dp(i-1,j) 在word1 的第i个位置新增 第i个字符得到, 因此 dp(i, j) = min(dp(i-1, j-1)+1, dp(i-1,j)+1, dp(i,j-1)+1)
// 若word1[i] == word2[j] 则dp(i,j) 取值还会是 dp(i-1, j-1), dp(i,j-1)+1, dp(i-1,j)+1 这三者之间的最小值，只不过dp(i-1,j-1)肯定会小于 dp(i,j-1)+1 和 dp(i-1,j)
// 因此当 word1[i] == word2[j] 时, dp(i,j) = dp(i-1, j-1)
// 考虑初始状态:
// dp(0, j) 相当于word1为空串时 与不为空串的word2 之间的距离, 因此dp(0, j) = j, 类似的 dp(i, 0) = i
// 所求的结果为  dp(i, j)
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}
	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j-1]+1, dp[i-1][j]+1, dp[i][j-1]+1)
			}
		}
	}
	return dp[m][n]
}
