package _004_longestOnes

// 1004. 最大连续1的个数 III
// 给定一个二进制数组 nums 和一个整数 k，假设最多可以翻转 k 个 0 ，则返回执行操作后 数组中连续 1 的最大个数 。
//
// 示例 1：
// 输入：nums = [1,1,1,0,0,0,1,1,1,1,0], K = 2
// 输出：6
// 解释：[1,1,1,0,0,1,1,1,1,1,1]
// 粗体数字从 0 翻转到 1，最长的子数组长度为 6。
//
// 示例 2：
// 输入：nums = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], K = 3
// 输出：10
// 解释：[0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
// 粗体数字从 0 翻转到 1，最长的子数组长度为 10。
//
// 提示：
//
// 1 <= nums.length <= 105
// nums[i] 不是 0 就是 1
// 0 <= k <= nums.length

func longestOnes(nums []int, k int) int {
	maxWith, N := 0, len(nums)
	swap := k
	for i, j := 0, 0; j < N; j++ {
		if nums[j] == 0 {
			if k == 0 {
				i = j + 1
				continue
			}
			if swap > 0 {
				swap = swap - 1
			} else {
				// i前移直到跳过遇到的第一个0
				for ; i <= j; i++ {
					if nums[i] == 0 {
						i = i + 1
						break
					}
				}
			}
		}
		maxWith = max(maxWith, j-i+1)
	}
	return maxWith
}
