package _7_isInterleave

// 给定三个字符串 s1、s2、s3，请你帮忙验证 s3 是否是由 s1 和 s2 交错 组成的。
//
// 两个字符串 s 和 t 交错 的定义与过程如下，其中每个字符串都会被分割成若干 非空 子字符串：
//
// s = s1 + s2 + ... + sn
// t = t1 + t2 + ... + tm
// |n - m| <= 1
// 交错 是 s1 + t1 + s2 + t2 + s3 + t3 + ... 或者 t1 + s1 + t2 + s2 + t3 + s3 + ...
// 注意：a + b 意味着字符串 a 和 b 连接。
//
// 示例 1：
//
// 输入：s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
// 输出：true
// 示例 2：
//
// 输入：s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
// 输出：false
// 示例 3：
//
// 输入：s1 = "", s2 = "", s3 = ""
// 输出：true
//
// 提示：
//
// 0 <= s1.length, s2.length <= 100
// 0 <= s3.length <= 200
// s1、s2、和 s3 都由小写英文字母组成
//
// 进阶：您能否仅使用 O(s2.length) 额外的内存空间来解决它?

// 动态规划解法 --TODO
// 设f(i,j) 代表 s1 的前i个元素 和 s2 的前j个元素 是否能交错组成 s3的前 i + j 个元素
// 若 s1[i] == s3[i+j]  且 f(i-1, j) 为true (即: s1 的前i+1 个元素和s2的前j个元素可以表示 s3的前i+j-1 个元素)
// 则 f(i,j) 也为true
// 类似的 有 s2[j] == s3[i+j] 且 f(i, j-1)为true, 则f(i,j) 也为true
// f(0,0) = true, f(1, 0) = (s1[0]==s3[0]), f(0,1) = (s2[0]==s3[0])
func isInterleave(s1 string, s2 string, s3 string) bool {
	l1, l2, l3 := len(s1), len(s2), len(s3)
	if l1+l2 != l3 {
		return false
	}

	dp := make([][]bool, l1+1)
	for i := range dp {
		dp[i] = make([]bool, l2+1)
	}

	dp[0][0] = true
	for i := 0; i <= l1; i++ {
		for j := 0; j <= l2; j++ {
			k := i + j - 1
			if i > 0 {
				dp[i][j] = dp[i][j] || (dp[i-1][j] && s1[i-1] == s3[k])
			}
			if j > 0 {
				dp[i][j] = dp[i][j] || (dp[i][j-1] && s2[j-1] == s3[k])
			}
		}
	}
	return dp[l1][l2]
}
