package _456_maxVowels

// 1456. 定长子串中元音的最大数目
// 给你字符串 s 和整数 k 。
//
// 请返回字符串 s 中长度为 k 的单个子字符串中可能包含的最大元音字母数。
//
// 英文中的 元音字母 为（a, e, i, o, u）。
//
// 示例 1：
// 输入：s = "abciiidef", k = 3
// 输出：3
// 解释：子字符串 "iii" 包含 3 个元音字母。
// 示例 2：
//
// 输入：s = "aeiou", k = 2
// 输出：2
// 解释：任意长度为 2 的子字符串都包含 2 个元音字母。
// 示例 3：
//
// 输入：s = "leetcode", k = 3
// 输出：2
// 解释："lee"、"eet" 和 "ode" 都包含 2 个元音字母。
// 示例 4：
//
// 输入：s = "rhythms", k = 4
// 输出：0
// 解释：字符串 s 中不含任何元音字母。
// 示例 5：
//
// 输入：s = "tryhard", k = 4
// 输出：1
//
// 提示：
// 1 <= s.length <= 10^5
// s 由小写英文字母组成
// 1 <= k <= s.length

// 超出时间限制
func maxVowels(s string, k int) int {
	maxNum, N := 0, len(s)
	vowels := map[byte]int{'a': 1, 'e': 1, 'i': 1, 'o': 1, 'u': 1}
	start, end := 0, min(N-1, k-1)
	for end < N {
		winCnt := 0
		for i := start; i <= end; i++ {
			ch := s[i]
			if vowels[ch] == 1 {
				winCnt += 1
			}
		}
		maxNum = max(maxNum, winCnt)
		start, end = start+1, end+1
	}
	return maxNum
}

// 使用滑动窗
func maxVowels1(s string, k int) int {
	winCnt, N := 0, len(s)
	vowels := map[byte]int{'a': 1, 'e': 1, 'i': 1, 'o': 1, 'u': 1}
	for i := 0; i < k; i++ {
		if vowels[s[i]] == 1 {
			winCnt += 1
		}
	}
	maxNum := winCnt
	for i, j := 1, k; j < N; i, j = i+1, j+1 {
		if vowels[s[i-1]] == 1 {
			winCnt -= 1
		}
		if vowels[s[j]] == 1 {
			winCnt += 1
		}
		maxNum = max(maxNum, winCnt)
	}
	return maxNum
}
