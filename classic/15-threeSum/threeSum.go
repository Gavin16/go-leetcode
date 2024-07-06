package _5_threeSum

import (
	"slices"
)

// 15. 三数之和
// 给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，
// 同时还满足 nums[i] + nums[j] + nums[k] == 0 。请
//
// 你返回所有和为 0 且不重复的三元组。
//
// 注意：答案中不可以包含重复的三元组。
//
// 示例 1：
//
// 输入：nums = [-1,0,1,2,-1,-4]
// 输出：[[-1,-1,2],[-1,0,1]]
// 解释：
// nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
// nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
// nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
// 不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
// 注意，输出的顺序和三元组的顺序并不重要。
// 示例 2：
//
// 输入：nums = [0,1,1]
// 输出：[]
// 解释：唯一可能的三元组和不为 0 。
// 示例 3：
//
// 输入：nums = [0,0,0]
// 输出：[[0,0,0]]
// 解释：唯一可能的三元组和为 0 。
//
// 提示：
//
// 3 <= nums.length <= 3000
// -105 <= nums[i] <= 105

// 对nums排序, 然后遍历nums, 对于每个位置 i, 相当于寻找 两数和为 -nums[i]
// 击败: 93.80%
func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	res := make([][]int, 0)

	for i := 0; i < len(nums)-2; {
		if nums[i] > 0 {
			break
		}
		left, right := i+1, len(nums)-1
		for left < right {
			if nums[left]+nums[right] > -nums[i] {
				right -= 1
			} else if nums[left]+nums[right] < -nums[i] {
				left += 1
			} else {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left, right = left+1, right-1
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			}
		}
		for i = i + 1; i < len(nums)-2; i++ {
			if nums[i] != nums[i-1] {
				break
			}
		}
	}
	return res
}
