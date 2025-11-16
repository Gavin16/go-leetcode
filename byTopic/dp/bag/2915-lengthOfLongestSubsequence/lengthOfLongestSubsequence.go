package _915_lengthOfLongestSubsequence

import "math"

// 2915. 和为目标值的最长子序列的长度
// 给你一个下标从 0 开始的整数数组 nums 和一个整数 target 。
//
// 返回和为 target 的 nums 子序列中，子序列 长度的最大值 。如果不存在和为 target 的子序列，返回 -1 。
//
// 子序列 指的是从原数组中删除一些或者不删除任何元素后，剩余元素保持原来的顺序构成的数组。
//
// 示例 1：
//
// 输入：nums = [1,2,3,4,5], target = 9
// 输出：3
// 解释：总共有 3 个子序列的和为 9 ：[4,5] ，[1,3,5] 和 [2,3,4] 。最长的子序列是 [1,3,5] 和 [2,3,4] 。所以答案为 3 。
// 示例 2：
//
// 输入：nums = [4,1,3,2,1,5], target = 7
// 输出：4
// 解释：总共有 5 个子序列的和为 7 ：[4,3] ，[4,1,2] ，[4,2,1] ，[1,1,5] 和 [1,3,2,1] 。最长子序列为 [1,3,2,1] 。所以答案为 4 。
// 示例 3：
//
// 输入：nums = [1,1,5,4,5], target = 3
// 输出：-1
// 解释：无法得到和为 3 的子序列。
//
// 提示：
//
// 1 <= nums.length <= 1000
// 1 <= nums[i] <= 1000
// 1 <= target <= 1000

func lengthOfLongestSubsequence(nums []int, target int) int {
	n := len(nums)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, target+1)
		for j := range memo[i] {
			memo[i][j] = -2
		}
	}
	var dfs func(i, left int) int
	dfs = func(i, left int) (ans int) {
		if i < 0 {
			if left == 0 {
				return 0
			}
			return -1
		}
		p := &memo[i][left]
		if *p != -2 {
			return *p
		}
		defer func() { *p = ans }()
		if left < nums[i] {
			return dfs(i-1, left)
		}
		skip := dfs(i-1, left)
		choose := dfs(i-1, left-nums[i])
		if choose != -1 {
			return max(choose+1, skip)
		}
		return skip
	}
	return dfs(n-1, target)
}

func lengthOfLongestSubsequence1(nums []int, target int) int {
	n := len(nums)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, target+1)
	}
	for j := 1; j <= target; j++ {
		f[0][j] = math.MinInt
	}
	for i, x := range nums {
		for j := range target + 1 {
			if j < x {
				f[i+1][j] = f[i][j]
			} else {
				f[i+1][j] = max(f[i][j], f[i][j-nums[i]]+1)
			}
		}
	}
	if f[n][target] > 0 {
		return f[n][target]
	}
	return -1
}
