package _2_generateParenthesis

// 22. 括号生成
// 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
//
// 示例 1：
//
// 输入：n = 3
// 输出：["((()))","(()())","(())()","()(())","()()()"]
// 示例 2：
//
// 输入：n = 1
// 输出：["()"]
//
// 提示：
//
// 1 <= n <= 8

// 击败: 100.00%
func generateParenthesis(n int) []string {

	ans := make([]string, 0)
	curr := make([]byte, 0)

	var dfs func(left, right int)
	dfs = func(left, right int) {
		if left == 0 && right == 0 {
			ans = append(ans, string(curr))
			return
		}

		if left > 0 {
			curr = append(curr, '(')
			dfs(left-1, right)
			curr = curr[:len(curr)-1]
		}

		if right > left {
			curr = append(curr, ')')
			dfs(left, right-1)
			curr = curr[:len(curr)-1]
		}
	}
	dfs(n, n)
	return ans
}
