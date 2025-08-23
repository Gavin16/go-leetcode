package _21_maximalSquare

import "testing"

func TestMaximalSquare(t *testing.T) {
	matrix1 := [][]byte{{0, 1}, {1, 0}}
	println(maximalSquare(matrix1))

	matrix2 := [][]byte{{0}}
	println(maximalSquare(matrix2))

	matrix3 := [][]byte{{1, 0, 1, 0, 0}, {1, 0, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 0, 0, 1, 0}}
	println(maximalSquare(matrix3))

}
