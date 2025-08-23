package _0_isValid

// 20. 有效的括号
// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
//
// 有效字符串需满足：
//
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 每个右括号都有一个对应的相同类型的左括号。
//
// 示例 1：
//
// 输入：s = "()"
// 输出：true
// 示例 2：
//
// 输入：s = "()[]{}"
// 输出：true
// 示例 3：
//
// 输入：s = "(]"
// 输出：false
//
// 提示：
//
// 1 <= s.length <= 104
// s 仅由括号 '()[]{}' 组成

// 击败:20.22%
func isValid(s string) bool {
	stack := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 {
				return false
			} else {
				left := stack[len(stack)-1]
				if max(left, s[i])-min(left, s[i]) > 2 {
					return false
				}
				stack = stack[:len(stack)-1]
			}
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}
