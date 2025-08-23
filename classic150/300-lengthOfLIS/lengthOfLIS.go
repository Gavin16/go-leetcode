package _00_lengthOfLIS

import "sort"

// 300. 最长递增子序列
// 给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
// 子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。
// 例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。
// 示例 1：
//
// 输入：nums = [10,9,2,5,3,7,101,18]
// 输出：4
// 解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。
// 示例 2：
//
// 输入：nums = [0,1,0,3,2,3]
// 输出：4
// 示例 3：
//
// 输入：nums = [7,7,7,7,7,7,7]
// 输出：1
//
// 提示：
//
// 1 <= nums.length <= 2500
// -104 <= nums[i] <= 104
//
// 进阶：
//
// 你能将算法的时间复杂度降低到 O(n log(n)) 吗?

// 深度优先搜索
// 当遇到第 i 个元素时
// 取或者不取该元素 存在两种路径
// 若 nums[i] > last  跳过则直接忽略，使用则直接算进来
// 若 nums[i] <= last 跳过则直接忽略，使用则需要一直清空序列中的后面大于等于 last 的元数

// 动态规划
// dp[i] 代表在 nums 数组以i位置结尾的最长递增序列长度
// 遍历 i 之前所有元素 j∈[0, i-1] 都与nums[i]比较, 若 nums[j] < nums[i] 则dp[i] = max(dp[i], dp[j]+1)
// 这样可以再 j 遍历范围(路径范围内) 找出i位置的最长递增序列的长度
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 1
	// 最长递增子序列可能出现在序列中的任何位置
	// 因此需要找出所有位置中最长的那个
	maxLen := 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
		maxLen = max(maxLen, dp[i])
	}
	return maxLen
}

// 使用 O(n log(n))  复杂度解法
func lengthOfLIS1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var sortedSeq = []int{}
	for _, n := range nums {
		idx := sort.Search(
			len(sortedSeq),
			func(i int) bool {
				return sortedSeq[i] >= n
			},
		)
		if idx == len(sortedSeq) {
			sortedSeq = append(sortedSeq, n)
		} else {
			sortedSeq[idx] = n
		}
	}
	return len(sortedSeq)
}

// 动态规划状态转移方程
// 定义dp[i] 为nums 中以位置i结尾 最长严格递增序列的长度
// 则 dp[i+1] 为 0~i 范围内的j,  dp[i+1] = max(dp[j] + 1, dp[i+1]) 对于nums[i+1] > nums[j]
// 其中dp[0] = 1, 遍历 j 前 dp[i+1] = 1
// 求nums 的最长严格递增子序列的长度, 需要找出dp数组中所有值的最大值
func lengthOfLIS2(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
	}
	maxLen := 1
	for _, n := range dp {
		maxLen = max(maxLen, n)
	}
	return maxLen
}
