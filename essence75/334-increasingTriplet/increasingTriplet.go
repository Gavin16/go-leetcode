package _34_increasingTriplet

import "math"

// 334. 递增的三元子序列
// 给你一个整数数组 nums ，判断这个数组中是否存在长度为 3 的递增子序列。
//
// 如果存在这样的三元组下标 (i, j, k) 且满足 i < j < k ，使得 nums[i] < nums[j] < nums[k] ，返回 true ；否则，返回 false 。
//
// 示例 1：
//
// 输入：nums = [1,2,3,4,5]
// 输出：true
// 解释：任何 i < j < k 的三元组都满足题意
// 示例 2：
//
// 输入：nums = [5,4,3,2,1]
// 输出：false
// 解释：不存在满足题意的三元组
// 示例 3：
//
// 输入：nums = [2,1,5,0,4,6]
// 输出：true
// 解释：其中一个满足题意的三元组是 (3, 4, 5)，因为 nums[3] == 0 < nums[4] == 4 < nums[5] == 6
//
// 提示：
//
// 1 <= nums.length <= 5 * 105
// -231 <= nums[i] <= 231 - 1
//
// 进阶：你能实现时间复杂度为 O(n) ，空间复杂度为 O(1) 的解决方案吗？

// 使用单调栈
func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	N := len(nums)
	left, right := make([]int, N), make([]int, N)
	left[0], right[N-1] = nums[0], nums[N-1]
	for i := 1; i < N; i++ {
		left[i] = min(left[i-1], nums[i])
	}
	for j := N - 2; j >= 0; j-- {
		right[j] = max(right[j+1], nums[j])
	}
	for k := 1; k < N-1; k++ {
		if nums[k] > left[k] && nums[k] < right[k] {
			return true
		}
	}
	return false
}

// 贪心解法
func increasingTriplet1(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	first, second := nums[0], math.MaxInt
	for i := 0; i < len(nums); i++ {
		if nums[i] > second {
			return true
		} else if nums[i] > first {
			second = nums[i]
		} else {
			first = nums[i]
		}
	}
	return false
}
