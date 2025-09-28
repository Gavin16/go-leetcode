package _071_gcdOfStrings

import "strings"

// 1071. 字符串的最大公因子
// 对于字符串 s 和 t，只有在 s = t + t + t + ... + t + t（t 自身连接 1 次或多次）时，我们才认定 “t 能除尽 s”。
// 给定两个字符串 str1 和 str2 。返回 最长字符串 x，要求满足 x 能除尽 str1 且 x 能除尽 str2 。
//
// 示例 1：
// 输入：str1 = "ABCABC", str2 = "ABC"
// 输出："ABC"
//
// 示例 2：
// 输入：str1 = "ABABAB", str2 = "ABAB"
// 输出："AB"
//
// 示例 3：
// 输入：str1 = "LEET", str2 = "CODE"
// 输出：""
//
// 提示：
// 1 <= str1.length, str2.length <= 1000
// str1 和 str2 由大写英文字母组成
func gcdOfStrings(str1 string, str2 string) string {
	if len(str2) > len(str1) {
		str1, str2 = str2, str1
	}
	for strings.HasPrefix(str1, str2) && len(str2) > 0 {
		str1, str2 = str2, strings.Replace(str1, str2, "", -1)
	}
	if len(str2) == 0 {
		return str1
	}
	return ""
}
