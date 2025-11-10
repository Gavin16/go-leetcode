package grid

import "math/rand"

func BuildRandomMatrix(m, n int, ceil int) [][]int {
	matrix := make([][]int, m)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j] = rand.Intn(ceil)
		}
	}
	return matrix
}
