package _4_largestRectangleArea

// 84. 柱状图中最大的矩形
// 给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
//
// 求在该柱状图中，能够勾勒出来的矩形的最大面积。
//
// 示例 1:
//
// 输入：heights = [2,1,5,6,2,3]
// 输出：10
// 解释：最大的矩形为图中红色区域，面积为 10
// 示例 2：
//
// 输入： heights = [2,4]
// 输出： 4
//
// 提示：
//
// 1 <= heights.length <=105
// 0 <= heights[i] <= 104

// 使用单调栈
// 找出height[i] 左侧
func largestRectangleArea(heights []int) int {
	n := len(heights)
	leftIdx := make([]int, n)
	rightIdx := make([]int, n)
	leftIdx[0], rightIdx[n-1] = -1, n

	for i := 1; i < n; i++ {
		temp := i - 1
		for temp >= 0 && heights[temp] >= heights[i] {
			temp = leftIdx[temp]
		}
		leftIdx[i] = temp
	}
	for j := n - 2; j >= 0; j-- {
		temp := j + 1
		for temp < n && heights[temp] >= heights[j] {
			temp = rightIdx[temp]
		}
		rightIdx[j] = temp
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans = max(ans, (rightIdx[i]-leftIdx[i]-1)*heights[i])
	}
	return ans
}
