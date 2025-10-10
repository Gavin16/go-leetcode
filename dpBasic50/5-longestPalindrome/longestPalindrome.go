package __longestPalindrome

// 5. 最长回文子串
// 给你一个字符串 s，找到 s 中最长的 回文 子串。
//
// 示例 1：
//
// 输入：s = "babad"
// 输出："bab"
// 解释："aba" 同样是符合题意的答案。
// 示例 2：
//
// 输入：s = "cbbd"
// 输出："bb"
//
// 提示：
//
// 1 <= s.length <= 1000
// s 仅由数字和英文字母组成

func longestPalindrome(s string) string {
	N := len(s)
	f := make([][]bool, N)
	for i := range f {
		f[i] = make([]bool, N)
		f[i][i] = true
	}
	ans := string(s[0])
	for i := N - 2; i >= 0; i-- {
		for j := i + 1; j < N; j++ {
			if j == i+1 {
				f[i][j] = s[i] == s[j]
			} else {
				f[i][j] = f[i+1][j-1] && (s[i] == s[j])
			}
			if f[i][j] && (j-i+1) > len(ans) {
				ans = s[i : j+1]
			}
		}
	}
	return ans
}
