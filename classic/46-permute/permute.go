package _6_permute

import "sort"

// 46. 全排列
// 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列。
// 你可以 按任意顺序 返回答案。
// 示例 1：
//
// 输入：nums = [1,2,3]
// 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
// 示例 2：
//
// 输入：nums = [0,1]
// 输出：[[0,1],[1,0]]
// 示例 3：
//
// 输入：nums = [1]
// 输出：[[1]]
//
// 提示：
// 1 <= nums.length <= 6
// -10 <= nums[i] <= 10
// nums 中的所有整数 互不相同

// 击败: 100.00%
func permute(nums []int) [][]int {
	ans := make([][]int, 0)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	var dfs func(i int)
	curr := make([]int, 0)
	dfs = func(i int) {
		if i == len(nums) {
			ans = append(ans, append([]int(nil), curr...))
			return
		}
		for k := i; k < len(nums); k++ {
			nums[i], nums[k] = nums[k], nums[i]
			curr = append(curr, nums[i])
			dfs(i + 1)
			curr = curr[:len(curr)-1]
			nums[i], nums[k] = nums[k], nums[i]
		}
	}
	dfs(0)
	return ans
}
