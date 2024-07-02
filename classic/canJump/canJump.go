package canJump

// 55. 跳跃游戏
// 给你一个非负整数数组 nums ，你最初位于数组的 第一个下标 。数组中的每个元素代表你在该位置可以跳跃的最大长度。
//
// 判断你是否能够到达最后一个下标，如果可以，返回 true ；否则，返回 false 。
// 示例 1：
//
// 输入：nums = [2,3,1,1,4]
// 输出：true
// 解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。
// 示例 2：
//
// 输入：nums = [3,2,1,0,4]
// 输出：false
// 解释：无论怎样，总会到达下标为 3 的位置。但该下标的最大跳跃长度是 0 ， 所以永远不可能到达最后一个下标。
//
//
// 提示：
//
// 1 <= nums.length <= 104
// 0 <= nums[i] <= 105

// 解法1: 回溯算法
// 每走一步都面临相同数量的同样的问题，至于最终位置是否可达 需要通过一直递归到最后才知道
// 这种情况正适合使用 深度优先搜索 进行处理
// 由于计算过程可能面临大量的重复情况，因此需要使用备忘录来避免重复计算
// 时间复杂度O(n), 空间复杂度O(n), 击败: 12.50%
func canJump(nums []int) bool {
	if len(nums) < 2 {
		return true
	}
	memo := make([]bool, len(nums))
	return dfsSearch(nums, memo, 0)
}

func dfsSearch(nums []int, memo []bool, pos int) bool {
	memo[pos] = true
	if pos+nums[pos] >= len(nums)-1 {
		return true
	} else {
		scope := nums[pos]
		canReach := false
		for i := 1; i <= scope; i++ {
			if memo[pos+i] == false {
				canReach = canReach || dfsSearch(nums, memo, pos+i)
			}
		}
		return canReach
	}
}

// 解法2
// 能跳跃最远距离为 maxJump, 遍历nums 一遍, 判断是否会出现 下标 i > maxJump
// 若出现 i > maxJump 说明已经跳不下去了； 否则能一直遍历到最后则代表可以跳跃
// 时间复杂度 O(n), 空间复杂度O(1) 击败: 88.74%
func canJump1(nums []int) bool {
	if len(nums) < 2 {
		return true
	}
	maxJump := 0
	for i, v := range nums {
		if i > maxJump {
			return false
		}
		maxJump = max(maxJump, i+v)
	}
	return true
}

// 解法3
// 遍历找出最远的长度
// 若最远长度小于 nums 则说明不能跳到最后，否则可以调到最后
// 时间复杂度O(n), 空间复杂度O(1) 击败: 94.33%
func canJump2(nums []int) bool {
	maxLen := 0

	for i := 0; i <= maxLen && maxLen < len(nums); i++ {
		if i+nums[i] > maxLen {
			maxLen = i + nums[i]
		}
	}
	return maxLen >= len(nums)-1
}
