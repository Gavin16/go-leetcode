package _7_combine

// 77. 组合
// 给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
// 你可以按 任何顺序 返回答案。
// 示例 1：
//
// 输入：n = 4, k = 2
// 输出：
// [
//
//	[2,4],
//	[3,4],
//	[2,3],
//	[1,2],
//	[1,3],
//	[1,4],
//
// ]
// 示例 2：
//
// 输入：n = 1, k = 1
// 输出：[[1]]
//
// 提示：
//
// 1 <= n <= 20
// 1 <= k <= n

// 击败: 73.17%
func combine(n int, k int) [][]int {
	ans := make([][]int, 0)
	curr := make([]int, 0)
	var dfs func(id, left int)
	dfs = func(id, left int) {
		if id > n || left == 0 {
			if len(curr) == k {
				ans = append(ans, append([]int(nil), curr...))
			}
			return
		}
		// 选择当前id
		curr = append(curr, id)
		dfs(id+1, left-1)
		curr = curr[:len(curr)-1]
		// 跳过当前id
		dfs(id+1, left)
	}
	dfs(1, k)
	return ans
}
