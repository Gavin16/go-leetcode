package _89_rotate

// 189. 轮转数组
// 给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。
//
// 示例 1:
//
// 输入: nums = [1,2,3,4,5,6,7], k = 3
// 输出: [5,6,7,1,2,3,4]
// 解释:
// 向右轮转 1 步: [7,1,2,3,4,5,6]
// 向右轮转 2 步: [6,7,1,2,3,4,5]
// 向右轮转 3 步: [5,6,7,1,2,3,4]
// 示例 2:
//
// 输入：nums = [-1,-100,3,99], k = 2
// 输出：[3,99,-1,-100]
// 解释:
// 向右轮转 1 步: [99,-1,-100,3]
// 向右轮转 2 步: [3,99,-1,-100]
//
//
// 提示：
//
// 1 <= nums.length <= 105
// -231 <= nums[i] <= 231 - 1
// 0 <= k <= 105
// 进阶：
//
// 尽可能想出更多的解决方案，至少有 三种 不同的方法可以解决这个问题。
// 你可以使用空间复杂度为 O(1) 的 原地 算法解决这个问题吗？

// 解法1: copy库函数操作切片 击败 69.33%, 时间O(n),空间 O(n)
func rotate(nums []int, k int) {
	k = k % len(nums)
	if k == 0 {
		return
	}
	part1 := make([]int, k)
	copy(part1, nums[len(nums)-k:])
	for i := len(nums) - k - 1; i >= 0; i-- {
		nums[i+k] = nums[i]
	}
	copy(nums[:k], part1)
}

// 解法2: 反转切片 29.23%   时间:O(n),空间O(1)
func rotate2(nums []int, k int) {
	k = k % len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}
func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

// 解法3:
func rotate3(nums []int, k int) {

}
