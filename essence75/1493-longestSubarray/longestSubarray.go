package _493_longestSubarray

// 1493. 删掉一个元素以后全为 1 的最长子数组
// 给你一个二进制数组 nums ，你需要从中删掉一个元素。
//
// 请你在删掉元素的结果数组中，返回最长的且只包含 1 的非空子数组的长度。
// 如果不存在这样的子数组，请返回 0 。
//
// 提示 1：
// 输入：nums = [1,1,0,1]
// 输出：3
// 解释：删掉位置 2 的数后，[1,1,1] 包含 3 个 1 。
//
// 示例 2：
// 输入：nums = [0,1,1,1,0,1,1,0,1]
// 输出：5
// 解释：删掉位置 4 的数字后，[0,1,1,1,1,1,0,1] 的最长全 1 子数组为 [1,1,1,1,1] 。
// 示例 3：
//
// 输入：nums = [1,1,1]
// 输出：2
// 解释：你必须要删除一个元素。
//
// 提示：
// 1 <= nums.length <= 105
// nums[i] 要么是 0 要么是 1

func longestSubarray(nums []int) int {
	left, right := 0, 0
	delFlag := false
	maxLen := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			if !delFlag {
				delFlag = true
				left, right = right, 0
			} else {
				left, right = right, 0
			}
		} else {
			right += 1
		}
		if !delFlag {
			maxLen = max(maxLen, right-1)
		} else {
			maxLen = max(maxLen, left+right)
		}
	}
	return maxLen
}
