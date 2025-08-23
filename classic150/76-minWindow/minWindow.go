package _6_minWindow

import (
	"math"
)

// 76. 最小覆盖子串
// 给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。
// 如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
//
// 注意：
//
// 对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
// 如果 s 中存在这样的子串，我们保证它是唯一的答案。
//
// 示例 1：
//
// 输入：s = "ADOBECODEBANC", t = "ABC"
// 输出："BANC"
// 解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。
// 示例 2：
//
// 输入：s = "a", t = "a"
// 输出："a"
// 解释：整个字符串 s 是最小覆盖子串。
// 示例 3:
//
// 输入: s = "a", t = "aa"
// 输出: ""
// 解释: t 中两个字符 'a' 均应包含在 s 的子串中，
// 因此没有符合条件的子字符串，返回空字符串。
//
// 提示：
//
// m == s.length
// n == t.length
// 1 <= m, n <= 105
// s 和 t 由英文字母组成
//
// 进阶：你能设计一个在 o(m+n) 时间内解决此问题的算法吗？

// 击败: 46.50%
func minWindow(s string, t string) string {
	cM1, cM2 := map[byte]int{}, map[byte]int{}
	notEq := map[byte]bool{}
	for i := 0; i < len(t); i++ {
		cM1[t[i]] = cM1[t[i]] + 1
		notEq[t[i]] = true
	}

	start, end := 0, 0
	minLen, ans := math.MaxInt32, ""
	for end < len(s) && start <= len(s)-len(t) {
		if len(notEq) != 0 {
			ch := s[end]
			if _, ok := cM1[ch]; ok {
				cM2[ch] += 1
				if cM2[ch] >= cM1[ch] {
					delete(notEq, ch)
				}
			}
			end += 1
		}
		for len(notEq) == 0 {
			currLen := end - start
			if currLen < minLen {
				minLen = currLen
				ans = s[start:end]
			}
			ch := s[start]
			if _, ok := cM1[ch]; ok {
				cM2[ch] -= 1
				if cM2[ch] < cM1[ch] {
					notEq[ch] = true
				}
			}
			start += 1
		}
	}
	return ans
}
