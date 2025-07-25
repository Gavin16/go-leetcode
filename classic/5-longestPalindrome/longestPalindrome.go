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
