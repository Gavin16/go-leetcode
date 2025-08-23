package _90_wordPattern

import "strings"

// 290. 单词规律
// 给定一种规律 pattern 和一个字符串 s ，判断 s 是否遵循相同的规律。
//
// 这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 s 中的每个非空单词之间存在着双向连接的对应规律。
//
// 示例1:
//
// 输入: pattern = "abba", s = "dog cat cat dog"
// 输出: true
// 示例 2:
//
// 输入:pattern = "abba", s = "dog cat cat fish"
// 输出: false
// 示例 3:
//
// 输入: pattern = "aaaa", s = "dog cat cat dog"
// 输出: false
//
// 提示:
//
// 1 <= pattern.length <= 300
// pattern 只包含小写英文字母
// 1 <= s.length <= 3000
// s 只包含小写英文字母和 ' '
// s 不包含 任何前导或尾随对空格
// s 中每个单词都被 单个空格 分隔

// 击败： 17.69%
func wordPattern(pattern string, s string) bool {
	split := strings.Split(s, " ")
	if len(split) != len(pattern) {
		return false
	}
	map1, map2 := map[byte]string{}, map[string]byte{}

	for i := 0; i < len(pattern); i++ {
		_, ok1 := map1[pattern[i]]
		_, ok2 := map2[split[i]]
		b, str := pattern[i], split[i]
		if !ok1 && !ok2 {
			map1[b] = str
			map2[str] = b
		} else if !(ok1 && ok2) {
			return false
		} else if map1[b] != str || map2[str] != b {
			return false
		}
	}
	return true
}
