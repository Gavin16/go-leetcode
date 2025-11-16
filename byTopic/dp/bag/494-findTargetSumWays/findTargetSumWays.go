package _94_findTargetSumWays

// 494. 目标和
// 给你一个非负整数数组 nums 和一个整数 target 。
//
// 向数组中的每个整数前添加 '+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：
// 例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。
// 返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。
//
// 示例 1：
// 输入：nums = [1,1,1,1,1], target = 3
// 输出：5
// 解释：一共有 5 种方法让最终目标和为 3 。
// -1 + 1 + 1 + 1 + 1 = 3
// +1 - 1 + 1 + 1 + 1 = 3
// +1 + 1 - 1 + 1 + 1 = 3
// +1 + 1 + 1 - 1 + 1 = 3
// +1 + 1 + 1 + 1 - 1 = 3
// 示例 2：
//
// 输入：nums = [1], target = 1
// 输出：1
//
// 提示：
//
// 1 <= nums.length <= 20
// 0 <= nums[i] <= 1000
// 0 <= sum(nums[i]) <= 1000
// -1000 <= target <= 1000

func findTargetSumWays(nums []int, target int) int {
	s, n := 0, len(nums)
	for _, v := range nums {
		s += v
	}
	s -= abs(target)
	if s < 0 || s%2 == 1 {
		return 0
	}
	t := s / 2
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, t+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(i, c int) int
	dfs = func(i, c int) (ans int) {
		if i == -1 {
			if c == 0 {
				return 1
			}
			return 0
		}
		p := &memo[i][c]
		if *p != -1 {
			return *p
		}
		defer func() { *p = ans }()
		if c < nums[i] {
			return dfs(i-1, c)
		}
		return dfs(i-1, c) + dfs(i-1, c-nums[i])
	}
	return dfs(n-1, t)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
