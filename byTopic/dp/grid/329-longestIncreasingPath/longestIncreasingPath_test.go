package _29_longestIncreasingPath

import (
	"fmt"
	"testing"
)

func TestLongestIncreasingPath(t *testing.T) {
	matrix := generateSpiralMatrix(200)
	for _, row := range matrix {
		fmt.Println(row)
	}
	tests := []struct {
		name  string
		input [][]int
		want  int
	}{
		{"1", [][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}}, 4},
		{"2", matrix, 40000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestIncreasingPath(tt.input)
			if got != tt.want {
				t.Errorf("longestIncreasingPath(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func generateSpiralMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	left, right, top, bottom := 0, n-1, 0, n-1
	num := 1

	for left <= right && top <= bottom {
		// → 从左到右
		for j := left; j <= right; j++ {
			matrix[top][j] = num
			num++
		}
		top++
		// ↓ 从上到下
		for i := top; i <= bottom; i++ {
			matrix[i][right] = num
			num++
		}
		right--
		// ← 从右到左
		if top <= bottom {
			for j := right; j >= left; j-- {
				matrix[bottom][j] = num
				num++
			}
			bottom--
		}
		// ↑ 从下到上
		if left <= right {
			for i := bottom; i >= top; i-- {
				matrix[i][left] = num
				num++
			}
			left++
		}
	}

	return matrix
}
