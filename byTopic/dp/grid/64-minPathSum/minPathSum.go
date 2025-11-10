package _4_minPathSum

import "math"

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	f := make([][]int, m)
	for i := range f {
		f[i] = make([]int, n)
	}
	f[0][0] = grid[0][0]
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			minSum := math.MaxInt
			if i > 0 {
				minSum = min(minSum, f[i-1][j])
			}
			if j > 0 {
				minSum = min(minSum, f[i][j-1])
			}
			f[i][j] = minSum + grid[i][j]
		}
	}
	return f[m-1][n-1]
}
