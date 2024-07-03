package _5_canJump2

// 45. 跳跃游戏 II
// 给定一个长度为 n 的 0 索引整数数组 nums。初始位置为 nums[0]。
//
// 每个元素 nums[i] 表示从索引 i 向前跳转的最大长度。换句话说，如果你在 nums[i] 处，你可以跳转到任意 nums[i + j] 处:
//
// 0 <= j <= nums[i]
// i + j < n
// 返回到达 nums[n - 1] 的最小跳跃次数。生成的测试用例可以到达 nums[n - 1]。
// 示例 1:
//
// 输入: nums = [2,3,1,1,4]
// 输出: 2
// 解释: 跳到最后一个位置的最小跳跃数是 2。
//
//	从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
//
// 示例 2:
//
// 输入: nums = [2,3,0,1,4]
// 输出: 2
//
// 提示:
//
// 1 <= nums.length <= 104
// 0 <= nums[i] <= 1000
// 题目保证可以到达 nums[n-1]

// 解法1 贪心算法
// 要求跳跃到最后位置的最小步数，这个跳跃过程中每一步都是最小部署，因此采用切面记录到达每个位置的最小步数
// 记录访问每个位置的最短跳跃次数
// 根据每一步能访问的最远位置,更新最短跳跃次数数组
// 击败: 13.53%
func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	minSteps := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		for j := 1; j <= nums[i] && i+j < len(nums); j++ {
			curr := i + j
			if minSteps[curr] == 0 {
				minSteps[curr] = minSteps[i] + 1
			} else {
				minSteps[curr] = min(minSteps[i]+1, minSteps[curr])
			}
		}
	}
	return minSteps[len(nums)-1]
}
