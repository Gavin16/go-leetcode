package _40_deleteAndEarn

// 740. 删除并获得点数
// 给你一个整数数组 nums ，你可以对它进行一些操作。
//
// 每次操作中，选择任意一个 nums[i] ，删除它并获得 nums[i] 的点数。
// 之后，你必须删除 所有 等于 nums[i] - 1 和 nums[i] + 1 的元素。
//
// 开始你拥有 0 个点数。返回你能通过这些操作获得的最大点数。
//
// 示例 1：
//
// 输入：nums = [3,4,2]
// 输出：6
// 解释：
// 删除 4 获得 4 个点数，因此 3 也被删除。
// 之后，删除 2 获得 2 个点数。总共获得 6 个点数。
// 示例 2：
//
// 输入：nums = [2,2,3,3,3,4]
// 输出：9
// 解释：
// 删除 3 获得 3 个点数，接着要删除两个 2 和 4 。
// 之后，再次删除 3 获得 3 个点数，再次删除 3 获得 3 个点数。
// 总共获得 9 个点数。
//
// 提示：
//
// 1 <= nums.length <= 2 * 104
// 1 <= nums[i] <= 104
func deleteAndEarn(nums []int) int {
	maxNum, cntMap := 0, map[int]int{}
	for _, num := range nums {
		maxNum = max(maxNum, num)
		cntMap[num] += 1
	}
	slice := make([]int, maxNum+1)
	for i := 1; i <= maxNum; i++ {
		if cnt, ok := cntMap[i]; ok {
			slice[i] = i * cnt
		}
	}
	if len(cntMap) == 1 {
		return nums[0] * cntMap[nums[0]]
	}

	pre, prePre := max(slice[0], slice[1]), slice[0]
	for i := 2; i < len(slice); i++ {
		curr := max(slice[i]+prePre, pre)
		pre, prePre = curr, pre
	}
	return pre
}
