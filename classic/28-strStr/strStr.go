package _8_strStr

// 28. 找出字符串中第一个匹配项的下标
// 给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串的第一个匹配项的下标（下标从 0 开始）。如果 needle 不是 haystack 的一部分，则返回  -1 。
//
// 示例 1：
//
// 输入：haystack = "sadbutsad", needle = "sad"
// 输出：0
// 解释："sad" 在下标 0 和 6 处匹配。
// 第一个匹配项的下标是 0 ，所以返回 0 。
// 示例 2：
//
// 输入：haystack = "leetcode", needle = "leeto"
// 输出：-1
// 解释："leeto" 没有在 "leetcode" 中出现，所以返回 -1 。
//
// 提示：
//
// 1 <= haystack.length, needle.length <= 104
// haystack 和 needle 仅由小写英文字符组成

// 击败: 100.00%
func strStr(haystack string, needle string) int {
	len1, len2, cnt := len(haystack), len(needle), 0
	for i, j := 0, 0; i < len1 && j < len2; {
		if haystack[i] == needle[j] {
			i, j = i+1, j+1
			cnt++
		} else {
			i, j = i-cnt+1, 0
			cnt = 0
		}
		if cnt == len2 {
			return i - len2
		}
	}
	return -1
}
