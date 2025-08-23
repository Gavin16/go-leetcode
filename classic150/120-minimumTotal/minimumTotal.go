package _20_minimumTotal

import "slices"

// 120. 三角形最小路径和
// 给定一个三角形 triangle ，找出自顶向下的最小路径和。
//
// 每一步只能移动到下一行中相邻的结点上。相邻的结点 在这里指的是 下标 与 上一层结点下标
// 相同或者等于 上一层结点下标 + 1 的两个结点。
// 就是说，如果正位于当前行的下标 i ，那么下一步可以移动到下一行的下标 i 或 i + 1 。
//
// 示例 1：
//
// 输入：triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
// 输出：11
// 解释：如下面简图所示：
//
//	 2
//	3 4
//
// 6 5 7
// 4 1 8 3
// 自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。
// 示例 2：
//
// 输入：triangle = [[-10]]
// 输出：-10
//
// 提示：
//
// 1 <= triangle.length <= 200
// triangle[0].length == 1
// triangle[i].length == triangle[i - 1].length + 1
// -104 <= triangle[i][j] <= 104
//
// 进阶：
// 你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题吗？

// 二维动态规划
// dp[i][j] 为三角形第i层，第j个位置作为路径终点  路径的最小和
// 从第i层到 第i+1层,
// 则 --若 0 < j < len(triangle[i+1]) - 1 :  dp[i+1][j] = min(dp[i][j-1]+triangle[i+1][j], dp[i][j]+triangle[i+1][j])
//
//	--若 j == 0 || j == len(triangle[i+1]) - 1 : dp[i+1][j] = dp[i][j] + triangle[i+1][j]
//
// dp[0][0] = triangle[0][0]
// 最后所求为 min(dp[i+1])

func minimumTotal(triangle [][]int) int {
	rowNum := len(triangle)
	dp := make([][]int, rowNum)
	for i := range dp {
		dp[i] = make([]int, rowNum)
	}
	dp[0][0] = triangle[0][0]

	for i := 1; i < rowNum; i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				dp[i][j] = dp[i-1][j] + triangle[i][j]
			} else if j == len(triangle[i])-1 {
				dp[i][j] = dp[i-1][j-1] + triangle[i][j]
			} else {
				dp[i][j] = min(dp[i-1][j-1]+triangle[i][j], dp[i-1][j]+triangle[i][j])
			}
		}
	}
	minTotal := slices.Min(dp[rowNum-1])
	return minTotal
}
