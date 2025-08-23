package _38_prodExceptSelf

// 238. 除自身以外数组的乘积
// 给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。
//
// 题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内。
//
// 请 不要使用除法，且在 O(n) 时间复杂度内完成此题。
// 示例 1:
//
// 输入: nums = [1,2,3,4]
// 输出: [24,12,8,6]
// 示例 2:
//
// 输入: nums = [-1,1,0,-3,3]
// 输出: [0,0,9,0,0]
//
// 提示：
//
// 2 <= nums.length <= 105
// -30 <= nums[i] <= 30
// 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内
//
// 进阶：你可以在 O(1) 的额外空间复杂度内完成这个题目吗？（出于对空间复杂度分析的目的，输出数组 不被视为 额外空间。）

// 计算所有乘积: 存在重复计算，时间复杂度O(n^2) 无法通过用例
func productExceptSelf(nums []int) []int {
	ret := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		part1 := nums[:i]
		part2 := nums[i+1:]
		ret[i] = product(part1) * product(part2)
	}
	return ret
}

func product(nums []int) int {
	result := 1
	for _, v := range nums {
		result = result * v
	}
	return result
}

// 预计算出各位置 左侧和右侧的乘积
// 时间复杂度 O(n) 空间复杂度 O(n) 击败: 76.29%
func productExceptSelf1(nums []int) []int {

	length := len(nums)
	result, left, right := make([]int, length), make([]int, length), make([]int, length)

	left[0] = 1
	for i := 1; i < len(nums); i++ {
		left[i] = left[i-1] * nums[i-1]
	}

	right[length-1] = 1
	for i := length - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}

	for i := 0; i < length; i++ {
		result[i] = left[i] * right[i]
	}
	return result
}

// 击败: 90.45%
func productExceptSelf2(nums []int) []int {

	length := len(nums)
	result := make([]int, length)

	// 先计算各位置左侧乘积列表
	result[0] = 1
	for i := 1; i < len(nums); i++ {
		result[i] = result[i-1] * nums[i-1]
	}

	// 动态计算右侧乘积列表
	rightProd := 1
	for i := length - 2; i >= 0; i-- {
		rightProd = rightProd * nums[i+1]
		result[i] = rightProd * result[i]
	}
	return result
}
