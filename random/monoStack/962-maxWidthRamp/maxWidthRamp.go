package _62_maxWidthRamp

// 962. 最大宽度坡
//
// 给定一个整数数组 A，坡是元组 (i, j)，其中  i < j 且 A[i] <= A[j]。这样的坡的宽度为 j - i。
//
// 找出 A 中的坡的最大宽度，如果不存在，返回 0 。
//
// 示例 1：
//
// 输入：[6,0,8,2,1,5]
// 输出：4
// 解释：
// 最大宽度的坡为 (i, j) = (1, 5): A[1] = 0 且 A[5] = 5.
// 示例 2：
//
// 输入：[9,8,1,0,1,9,4,0,4,1]
// 输出：7
// 解释：
// 最大宽度的坡为 (i, j) = (2, 9): A[2] = 1 且 A[9] = 1.
//
// 提示：
//
// 2 <= A.length <= 50000
// 0 <= A[i] <= 50000

func maxWidthRamp(nums []int) int {
	//return tree(nums, 0, len(nums)-1)
	mono := make([]int, 0)
	mono = append(mono, 0)
	maxWith := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] >= nums[mono[len(mono)-1]] {
			for j := len(mono) - 1; j >= 0 && nums[i] >= nums[mono[j]]; j-- {
				maxWith = max(maxWith, i-mono[j])
			}
		} else {
			mono = append(mono, i)
		}
	}
	return maxWith
}

//func tree(nums []int, left, right int) int {
//	if left >= right {
//		return 0
//	}
//	if nums[right] >= nums[left] {
//		return right - left
//	} else {
//		leftMax := tree(nums, left+1, right)
//		rightMax := tree(nums, left, right-1)
//		return max(leftMax, rightMax)
//	}
//}
