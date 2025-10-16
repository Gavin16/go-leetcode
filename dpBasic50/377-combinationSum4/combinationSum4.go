package _77_combinationSum4

import "sort"

// 377. 组合总和 Ⅳ
// 给你一个由 不同 整数组成的数组 nums ，和一个目标整数 target 。请你从 nums 中找出并返回总和为 target 的元素组合的个数。
//
// 题目数据保证答案符合 32 位整数范围。
//
// 示例 1：
//
// 输入：nums = [1,2,3], target = 4
// 输出：7
// 解释：
// 所有可能的组合为：
// (1, 1, 1, 1)
// (1, 1, 2)
// (1, 2, 1)
// (1, 3)
// (2, 1, 1)
// (2, 2)
// (3, 1)
// 请注意，顺序不同的序列被视作不同的组合。
// 示例 2：
//
// 输入：nums = [9], target = 3
// 输出：0
//
// 提示：
//
// 1 <= nums.length <= 200
// 1 <= nums[i] <= 1000
// nums 中的所有元素 互不相同
// 1 <= target <= 1000

func combinationSum4(nums []int, target int) int {
	N := len(nums)
	sort.Ints(nums)
	memo := make([][]int, N)
	for i := range memo {
		memo[i] = make([]int, target+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(i, left int) int
	dfs = func(i, left int) (ans int) {
		if left == 0 {
			return 1
		}
		p := &memo[i][left]
		if *p != -1 {
			return *p
		}
		defer func() { *p = ans }()
		cnt := 0
		for j := 0; j < N; j++ {
			if left >= nums[j] {
				cnt += dfs(j, left-nums[j])
			}
		}
		return cnt
	}
	return dfs(0, target)
}
