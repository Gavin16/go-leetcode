package _18_maxSubarraySumCircular

// 918. 环形子数组的最大和
// 给定一个长度为 n 的环形整数数组 nums ，返回 nums 的非空 子数组 的最大可能和 。
//
// 环形数组 意味着数组的末端将会与开头相连呈环状。形式上， nums[i] 的下一个元素是 nums[(i + 1) % n] ，
// nums[i] 的前一个元素是 nums[(i - 1 + n) % n] 。
// 子数组 最多只能包含固定缓冲区 nums 中的每个元素一次。形式上，对于子数组 nums[i], nums[i + 1], ..., nums[j] ，
// 不存在 i <= k1, k2 <= j 其中 k1 % n == k2 % n 。
//
// 示例 1：
// 输入：nums = [1,-2,3,-2]
// 输出：3
// 解释：从子数组 [3] 得到最大和 3
// 示例 2：
//
// 输入：nums = [5,-3,5]
// 输出：10
// 解释：从子数组 [5,5] 得到最大和 5 + 5 = 10
// 示例 3：
//
// 输入：nums = [3,-2,2,-3]
// 输出：3
// 解释：从子数组 [3] 和 [3,-2,2] 都可以得到最大和 3
//
// 提示：
//
// n == nums.length
// 1 <= n <= 3 * 104
// -3 * 104 <= nums[i] <= 3 * 104

// O(n^2) 时间复杂度
func maxSubarraySumCircular(nums []int) int {
	maxSum := nums[0]
	l := len(nums)
	for i := 0; i < l; i++ {
		currMax := nums[i]
		for j := i + 1; j < i+l; j++ {
			currMax = max(nums[j%l], currMax+nums[j%l])
			maxSum = max(maxSum, currMax)
		}
	}
	return maxSum
}

// 动态规划求解
// 记录数组所有从左到右 以下标 i 结束的子数组的最大和为  leftMax[i]
// 求循环数组的最大和 等价于 左子数组和 与 右子数组和的 最大值
// 定义 leftMax[i] 保存最长以 i 结尾的最大子数组和 则有
// 状态转移方程: leftMax[i] = max(leftMax[i-1], leftSum) 其中leftSum 为0~i 范围内所有元素的和
// 循环数组的最大子数组和 可以分为两种情况
// 1. 全部来自左子数组 --> 直接采用 kadane 算法即可实现
// 2. 最大子数组来自 左边的左子数组 和 右边的右子数组 两部分
//        这种情况有一些特点: 假设 左边数组部分为 0~i 右边数组部分为 j ~ n-1 其中 i <= j
//        a. 右边的右子数组需要时全部连续的, 因为它需要跨越数组末尾一直延伸到左子数组部分
//        b. 而左边最大和子数组的长度，视情况 最小为1, 最大可能为 j 也就是 i==j-1
//        c. 不管左子数组的长度是多少，以 i 结尾 子数组和最大的值 就是 从左到右以下标0开头子数组累加和的最大值 即
// 					max(leftMax[i-1], leftSum)
//        d. 这里以下标0开头 十分关键, 因为这种情况就是横跨了左右子数组两部分
// 所以leftMax 的计算方式为 leftMax[i] = max(leftMax[i-1], leftSum)

func maxSubarraySumCircular1(nums []int) int {
	leftMax := make([]int, len(nums))
	leftSum, res := nums[0], nums[0]
	leftMax[0] = nums[0]
	currMax := nums[0]
	n := len(nums)
	for i := 1; i < len(nums); i++ {
		currMax = max(nums[i], currMax+nums[i])
		res = max(res, currMax)

		leftSum += nums[i]
		leftMax[i] = max(leftMax[i-1], leftSum)
	}
	rightSum := 0
	for k := n - 1; k > 0; k-- {
		rightSum += nums[k]
		res = max(res, leftMax[k-1]+rightSum)
	}
	return res
}
