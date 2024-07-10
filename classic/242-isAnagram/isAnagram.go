package _42_isAnagram

// 242. 有效的字母异位词
// 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
//
// 注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。
//
// 示例 1:
//
// 输入: s = "anagram", t = "nagaram"
// 输出: true
// 示例 2:
//
// 输入: s = "rat", t = "car"
// 输出: false
//
// 提示:
//
// 1 <= s.length, t.length <= 5 * 104
// s 和 t 仅包含小写字母
//
// 进阶: 如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？

// 击败: 100.00%
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sMap, tMap := [26]int{}, [26]int{}
	for i := 0; i < len(s); i++ {
		sMap[s[i]-'a'] += 1
		tMap[t[i]-'a'] += 1
	}
	for i := 0; i < 26; i++ {
		if sMap[i] != tMap[i] {
			return false
		}
	}
	return true
}
