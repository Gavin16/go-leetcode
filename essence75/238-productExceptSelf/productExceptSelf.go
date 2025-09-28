package _38_productExceptSelf

// 238. 除自身以外数组的乘积
// 给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。
//
// 题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内。
//
// 请 不要使用除法，且在 O(n) 时间复杂度内完成此题。
//
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
// 输入 保证 数组 answer[i] 在  32 位 整数范围内
//
// 进阶：你可以在 O(1) 的额外空间复杂度内完成这个题目吗？（ 出于对空间复杂度分析的目的，输出数组 不被视为 额外空间。）

func productExceptSelf(nums []int) []int {
	mul, N := 1, len(nums)
	ans := make([]int, N)
	zeroCnt := 0
	for i := 0; i < N; i++ {
		if nums[i] == 0 {
			zeroCnt += 1
		} else {
			mul *= nums[i]
		}
	}
	if zeroCnt >= 2 {
		return make([]int, N)
	} else {
		for j := 0; j < N; j++ {
			if zeroCnt > 0 {
				if nums[j] == 0 {
					ans[j] = mul
				} else {
					ans[j] = 0
				}
			} else {
				ans[j] = mul / nums[j]
			}
		}
		return ans
	}
}
