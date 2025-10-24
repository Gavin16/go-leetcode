package backtrace

import "math"

// 132. 分割回文串 II
// 给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是回文串。
// 返回符合要求的 最少分割次数 。
//
// 示例 1：
//
// 输入：s = "aab"
// 输出：1
// 解释：只需一次分割就可将 s 分割成 ["aa","b"] 这样两个回文子串。
// 示例 2：
//
// 输入：s = "a"
// 输出：0
// 示例 3：
//
// 输入：s = "ab"
// 输出：1
//
// 提示：
//
// 1 <= s.length <= 2000
// s 仅由小写英文字母组成

func minCut(s string) int {
	N := len(s)
	memo := make([]int, N+1)
	for i := range memo {
		memo[i] = -1
	}
	var dfs func(n int) int
	dfs = func(n int) int {
		if isSymmetrical(s[:n+1]) {
			return 0
		}
		p := &memo[n]
		if *p != -1 {
			return *p
		}
		minCnt := 1999
		for i := 1; i <= n; i++ {
			if isSymmetrical(s[i : n+1]) {
				minCnt = min(minCnt, dfs(i-1)+1)
			}
		}
		*p = minCnt
		return minCnt
	}
	return dfs(N - 1)
}

func minCut1(s string) int {
	N := len(s)
	dp := make([]int, N)
	for i := 1; i < N; i++ {
		minCnt := math.MaxInt
		if isSymmetrical(s[:i+1]) {
			dp[i] = 0
		} else {
			for j := 1; j <= i; j++ {
				if isSymmetrical(s[j : i+1]) {
					minCnt = min(minCnt, dp[j-1]+1)
				}
			}
			dp[i] = minCnt
		}
	}
	return dp[N-1]
}

func isSymmetrical(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
