package _393_countPathsWithXorValue

import (
	"math/bits"
	"slices"
)

// 给你一个大小为 m x n 的二维整数数组 grid 和一个整数 k 。
//
// 你的任务是统计满足以下 条件 且从左上格子 (0, 0) 出发到达右下格子 (m - 1, n - 1) 的路径数目：
//
// 每一步你可以向右或者向下走，也就是如果格子存在的话，可以从格子 (i, j) 走到格子 (i, j + 1) 或者格子 (i + 1, j) 。
// 路径上经过的所有数字 XOR 异或值必须 等于 k 。
// 请你返回满足上述条件的路径总数。
//
// 由于答案可能很大，请你将答案对 109 + 7 取余 后返回。
//
// 示例 1：
//
// 输入：grid = [[2, 1, 5], [7, 10, 0], [12, 6, 4]], k = 11
//
// 输出：3
//
// 解释：
//
// 3 条路径分别为：
//
// (0, 0) → (1, 0) → (2, 0) → (2, 1) → (2, 2)
// (0, 0) → (1, 0) → (1, 1) → (1, 2) → (2, 2)
// (0, 0) → (0, 1) → (1, 1) → (2, 1) → (2, 2)
// 示例 2：
//
// 输入：grid = [[1, 3, 3, 3], [0, 3, 3, 2], [3, 0, 1, 1]], k = 2
//
// 输出：5
//
// 解释：
//
// 5 条路径分别为：
//
// (0, 0) → (1, 0) → (2, 0) → (2, 1) → (2, 2) → (2, 3)
// (0, 0) → (1, 0) → (1, 1) → (2, 1) → (2, 2) → (2, 3)
// (0, 0) → (1, 0) → (1, 1) → (1, 2) → (1, 3) → (2, 3)
// (0, 0) → (0, 1) → (1, 1) → (1, 2) → (2, 2) → (2, 3)
// (0, 0) → (0, 1) → (0, 2) → (1, 2) → (2, 2) → (2, 3)
// 示例 3：
//
// 输入：grid = [[1, 1, 1, 2], [3, 0, 3, 2], [3, 0, 2, 2]], k = 10
//
// 输出：0
//
// 提示：
//
// 1 <= m == grid.length <= 300
// 1 <= n == grid[r].length <= 300
// 0 <= grid[r][c] < 16
// 0 <= k < 16

func countPathsWithXorValue(grid [][]int, k int) int {
	mod := 1000_000_007
	m, n := len(grid), len(grid[0])
	maxVal := 0
	for _, row := range grid {
		maxVal = max(maxVal, slices.Max(row))
	}
	ceil := 1 << bits.Len(uint(maxVal))
	if k >= ceil {
		return 0
	}
	memo := make([][][]int, m)
	for i := range memo {
		memo[i] = make([][]int, n)
		for j := range memo[i] {
			memo[i][j] = make([]int, ceil)
			for t := range memo[i][j] {
				memo[i][j][t] = -1
			}
		}
	}
	var dfs func(sum int, r, c int) int
	dfs = func(sum int, r, c int) int {
		if r < 0 || c < 0 {
			return 0
		}
		val := grid[r][c]
		if r == 0 && c == 0 {
			if sum == val {
				return 1
			}
			return 0
		}
		p := &memo[r][c][sum]
		if *p != -1 {
			return *p
		}
		left := dfs(sum^grid[r][c], r, c-1) % mod
		up := dfs(sum^grid[r][c], r-1, c) % mod
		*p = left + up
		return (left + up) % mod
	}
	return dfs(k, m-1, n-1)
}

func countPathsWithXorValue1(grid [][]int, k int) int {
	// 改DP写法, 在递归调用回的阶段, 使用数组保存状态
	mod := 1000_000_007
	m, n := len(grid), len(grid[0])
	dp := make([][][]int, m+1)
	maxVal := 0
	for _, row := range grid {
		maxVal = max(maxVal, slices.Max(row))
	}
	ceil := 1 << bits.Len(uint(maxVal))
	if k >= ceil {
		return 0
	}
	for idx := range dp {
		dp[idx] = make([][]int, n+1)
		for j := range dp[idx] {
			dp[idx][j] = make([]int, ceil)
		}
	}
	dp[1][1][grid[0][0]] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for s := 0; s < ceil; s++ {
				if i == 0 && j == 0 && s == grid[i][j] {
					continue
				}
				dp[i+1][j+1][s] = (dp[i][j+1][s^grid[i][j]] + dp[i+1][j][s^grid[i][j]]) % mod
			}
		}
	}
	return dp[m][n][k]
}
