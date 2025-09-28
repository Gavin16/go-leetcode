package _16_combinationSum3

// 216. 组合总和 III
// 找出所有相加之和为 n 的 k 个数的组合，且满足下列条件：
//
// 只使用数字1到9
// 每个数字 最多使用一次
// 返回 所有可能的有效组合的列表 。该列表不能包含相同的组合两次，组合可以以任何顺序返回。
//
// 示例 1:
//
// 输入: k = 3, n = 7
// 输出: [[1,2,4]]
// 解释:
// 1 + 2 + 4 = 7
// 没有其他符合的组合了。
// 示例 2:
//
// 输入: k = 3, n = 9
// 输出: [[1,2,6], [1,3,5], [2,3,4]]
// 解释:
// 1 + 2 + 6 = 9
// 1 + 3 + 5 = 9
// 2 + 3 + 4 = 9
// 没有其他符合的组合了。
// 示例 3:
//
// 输入: k = 4, n = 1
// 输出: []
// 解释: 不存在有效的组合。
// 在[1,9]范围内使用4个不同的数字，我们可以得到的最小和是1+2+3+4 = 10，因为10 > 1，没有有效的组合。
//
// 提示:
//
// 2 <= k <= 9
// 1 <= n <= 60

func combinationSum3(k int, n int) [][]int {
	minSum := (k * (k + 1)) / 2
	if minSum > n {
		return [][]int{}
	}
	dict := make([]int, 10)
	var ans [][]int
	var dfs func(cnt, sum, target int, start int)
	dfs = func(cnt, sum, target int, start int) {
		if sum > target {
			return
		}
		if cnt == k {
			if sum == target {
				comb := make([]int, 0)
				for i := 1; i <= 9; i++ {
					if dict[i] == 1 {
						comb = append(comb, i)
					}
				}
				ans = append(ans, comb)
			}
			return
		}
		for i := start; i <= 9; i++ {
			dict[i] = 1
			dfs(cnt+1, sum+i, target, i+1)
			dict[i] = 0
		}
	}
	dfs(0, 0, n, 1)
	return ans
}

func combinationSum(k int, n int) [][]int {
	var temp []int
	var ans [][]int
	var dfs func(i, rest int)
	dfs = func(i, rest int) {
		if len(temp) == k && rest == 0 {
			ans = append(ans, append([]int{}, temp...))
			return
		}
		if len(temp)+10-i < k || rest < 0 {
			return
		}

		temp = append(temp, i)
		dfs(i+1, rest-i)
		temp = temp[:len(temp)-1]
		dfs(i+1, rest)
	}
	dfs(1, n)
	return ans
}
