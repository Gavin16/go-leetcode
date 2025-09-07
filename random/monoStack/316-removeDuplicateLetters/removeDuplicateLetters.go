package _16_removeDuplicateLetters

// 316. 去除重复字母
// 给你一个字符串 s ，请你去除字符串中重复的字母，使得每个字母只出现一次。
// 需保证 返回结果的字典序最小（要求不能打乱其他字符的相对位置）。
//
// 示例 1：
//
// 输入：s = "bcabc"
// 输出："abc"
// 示例 2：
//
// 输入：s = "cbacdcbc"
// 输出："acdb"
//
// 提示：
//
// 1 <= s.length <= 104
// s 由小写英文字母组成

// 使用单调栈保存需要保留的字符, 栈中元数单调递增
// 使用left 数组记录剩余字符数
// 使用inStack记录字符是否在栈中,避免每次判断时都要遍历
func removeDuplicateLetters(s string) string {
	stack := make([]byte, 0)
	inStack := [26]bool{}
	left := [26]int{}

	for _, ch := range s {
		left[ch-'a'] += 1
	}
	for id := range s {
		ch := s[id]
		if !inStack[ch-'a'] {
			for len(stack) > 0 && ch < stack[len(stack)-1] {
				last := stack[len(stack)-1]
				if left[last-'a'] == 0 {
					break
				}
				stack = stack[:len(stack)-1]
				inStack[last-'a'] = false
			}
			stack = append(stack, ch)
			inStack[ch-'a'] = true
		}
		left[ch-'a'] -= 1
	}
	return string(stack)
}
