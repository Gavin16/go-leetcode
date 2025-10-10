package _73_findNumberOfLIS

// 673. 最长递增子序列的个数
// 给定一个未排序的整数数组 nums ， 返回最长递增子序列的个数 。
//
// 注意 这个数列必须是 严格 递增的。
//
// 示例 1:
//
// 输入: [1,3,5,4,7]
// 输出: 2
// 解释: 有两个最长递增子序列，分别是 [1, 3, 4, 7] 和[1, 3, 5, 7]。
// 示例 2:
//
// 输入: [2,2,2,2,2]
// 输出: 5
// 解释: 最长递增子序列的长度是1，并且存在5个子序列的长度为1，因此输出5。
//
// 提示:
//
// 1 <= nums.length <= 2000
// -106 <= nums[i] <= 106

// f(i) 代表以nums[i] 结尾的最长递增子序列的长度
// 若 nums[i] > nums[j] --> f(i) = max(f(i), f(j))
// 若 nums[i] < nums[j] --> f(i) = max(f(i), f(j)+1)
func findNumberOfLIS(nums []int) int {
	N := len(nums)
	dp := make([]int, N)
	cnt := make([]int, N)
	ans, maxLen := 1, 0
	for i, x := range nums {
		dp[i], cnt[i] = 1, 1
		for j, y := range nums[:i] {
			if y < x {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
					cnt[i] = cnt[j]
				} else if dp[j]+1 == dp[i] {
					cnt[i] += cnt[j]
				}
			}
		}
		if dp[i] > maxLen {
			maxLen = dp[i]
			ans = cnt[i]
		} else if dp[i] == maxLen {
			ans += cnt[i]
		}
	}
	return ans
}
