package _289_minFallingPathSum

import "math"

// 1289. 下降路径最小和 II
// 给你一个 n x n 整数矩阵 grid ，请你返回 非零偏移下降路径 数字和的最小值。
//
// 非零偏移下降路径 定义为：从 grid 数组中的每一行选择一个数字，且按顺序选出来的数字中，相邻数字不在原数组的同一列。
//
// 示例 1：
//
// 输入：grid = [[1,2,3],[4,5,6],[7,8,9]]
// 输出：13
// 解释：
// 所有非零偏移下降路径包括：
// [1,5,9], [1,5,7], [1,6,7], [1,6,8],
// [2,4,8], [2,4,9], [2,6,7], [2,6,8],
// [3,4,8], [3,4,9], [3,5,7], [3,5,9]
// 下降路径中数字和最小的是 [1,5,7] ，所以答案是 13
// 示例 2：
//
// 输入：grid = [[7]]
// 输出：7
//
// 提示：
//
// n == grid.length == grid[i].length
// 1 <= n <= 200
// -99 <= grid[i][j] <= 99

func minFallingPathSum(grid [][]int) int {
	n := len(grid)
	f := make([][]int, n)
	for id := range f {
		f[id] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		f[0][i] = grid[0][i]
	}

	for i := 1; i < n; i++ {
		for j := 0; j < n; j++ {
			var left, right int
			if j > 0 && j < n-1 {
				left = findMin(f[i-1][0:j])
				right = findMin(f[i-1][j+1:])
			} else {
				if j == 0 {
					left = math.MaxInt
					right = findMin(f[i-1][j+1:])
				} else {
					left = findMin(f[i-1][0:j])
					right = math.MaxInt
				}
			}
			f[i][j] = min(left, right) + grid[i][j]
		}
	}
	ans := math.MaxInt
	for i := 0; i < n; i++ {
		ans = min(ans, f[n-1][i])
	}
	return ans
}

func findMin(arr []int) int {
	minVal := math.MaxInt
	for _, num := range arr {
		minVal = min(minVal, num)
	}
	return minVal
}
