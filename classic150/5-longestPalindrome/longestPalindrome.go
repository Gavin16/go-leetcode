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

// 暴力解法
func longestPalindrome(s string) string {
	l := len(s)
	f := make([][]int, l)
	for id := range f {
		f[id] = make([]int, l)
	}
	for k := 0; k < l; k++ {
		f[k][k] = 1
	}

	maxLen, maxLeft, maxRight := 1, 0, 0
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			isOK := isPalindrome(s, i, j)
			if isOK {
				f[i][j] = j - i + 1
			} else {
				f[i][j] = f[i][j-1]
			}

			if f[i][j] > maxLen {
				maxLen = f[i][j]
				maxLeft, maxRight = i, j
			}
		}
	}
	return s[maxLeft : maxRight+1]
}

// 判断s 的子串是否是回文字符串
func isPalindrome(s string, start, end int) bool {
	for end >= start {
		if s[start] != s[end] {
			return false
		} else {
			start += 1
			end -= 1
		}
	}
	return true
}

// 动态规划解法
// (1) 定义f(i,j) 表示s 子串s[i,j] 是否可以表示为回文字符串
// (2) f(i,j+1) = f(i+1, j) && s[i] == s[j+1]
// (3) 对于 i==j f(i,j) = true
//
//	对于 i+1=j f(i,j) = s[i] == s[j]
//
// (4) 对于
func longestPalindrome1(s string) string {
	l := len(s)
	f := make([][]bool, l)
	for id := range f {
		f[id] = make([]bool, l)
	}
	for k := 0; k < l; k++ {
		f[k][k] = true
	}

	maxLen, left, right := 1, 0, 0
	for i := l - 1; i >= 0; i-- {
		for j := i + 1; j < l; j++ {
			k := j - i + 1
			if k == 2 {
				f[i][j] = s[i] == s[j]
			} else {
				f[i][j] = f[i+1][j-1] && (s[i] == s[j])
			}
			if f[i][j] && k > maxLen {
				maxLen = k
				left, right = i, j
			}
		}
	}
	return s[left : right+1]
}
