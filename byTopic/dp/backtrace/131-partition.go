package backtrace

import "slices"

// 131. 分割回文串
// 给你一个字符串 s，请你将 s 分割成一些 子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。
//
// 示例 1：
//
// 输入：s = "aab"
// 输出：[["a","a","b"],["aa","b"]]
// 示例 2：
//
// 输入：s = "a"
// 输出：[["a"]]
//
// 提示：
//
// 1 <= s.length <= 16
// s 仅由小写英文字母组成

func partition(s string) [][]string {
	ans := make([][]string, 0)
	path := make([]string, 0)
	var dfs func(s string)
	dfs = func(s string) {
		n := len(s)
		if n == 0 {
			ans = append(ans, slices.Clone(path))
			return
		}
		for i := 1; i <= n; i++ {
			subS := s[:i]
			if isPalindrome(subS) {
				path = append(path, subS)
				dfs(s[i:])
				path = path[:len(path)-1]
			}
		}
	}
	dfs(s)
	return ans
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
