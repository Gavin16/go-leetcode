package _2_trap

// 42 接雨水
// 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
//
// 示例 1：
// 输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
// 输出：6
// 解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
//
// 示例 2：
//
// 输入：height = [4,2,0,3,2,5]
// 输出：9
//
// 提示：
//
// n == height.length
// 1 <= n <= 2 * 104
// 0 <= height[i] <= 105

// 击败 39.96%
func trap(height []int) int {
	maxLeft, sum := 0, 0
	// 正向求雨水
	for i := 1; i < len(height); i++ {
		if height[i] >= height[maxLeft] {
			top := min(height[i], height[maxLeft])
			for j := maxLeft + 1; j < i; j++ {
				sum += top - height[j]
			}
			maxLeft = i
		}
	}
	// 反向求雨水
	maxRight := len(height) - 1
	for k := len(height) - 2; k >= maxLeft; k-- {
		if height[k] >= height[maxRight] {
			top := min(height[k], height[maxRight])
			for j := k + 1; j < maxRight; j++ {
				sum += top - height[j]
			}
			maxRight = k
		}
	}
	return sum
}
