package _05_isIsomorphic

// 205. 同构字符串
// 给定两个字符串 s 和 t ，判断它们是否是同构的。
//
// 如果 s 中的字符可以按某种映射关系替换得到 t ，那么这两个字符串是同构的。
//
// 每个出现的字符都应当映射到另一个字符，同时不改变字符的顺序。不同字符不能映射到同一个字符上，相同字符只能映射到同一个字符上，字符可以映射到自己本身。
//
// 示例 1:
//
// 输入：s = "egg", t = "add"
// 输出：true
// 示例 2：
//
// 输入：s = "foo", t = "bar"
// 输出：false
// 示例 3：
//
// 输入：s = "paper", t = "title"
// 输出：true
//
// 提示：
//
// 1 <= s.length <= 5 * 104
// t.length == s.length
// s 和 t 由任意有效的 ASCII 字符组成

// 击败: 23.88%
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m1 := map[byte]byte{}
	m2 := map[byte]byte{}

	for i := 0; i < len(s); i++ {
		_, ok1 := m1[s[i]]
		_, ok2 := m2[t[i]]
		if !ok1 && !ok2 {
			m1[s[i]] = t[i]
			m2[t[i]] = s[i]
		} else if !(ok1 && ok2) {
			return false
		} else if !(m1[s[i]] == t[i] && m2[t[i]] == s[i]) {
			return false
		}
	}
	return true
}
