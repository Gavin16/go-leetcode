package _4_searchMatrix

import "testing"

func TestSearchMatrix(t *testing.T) {
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	target := 3
	println(searchMatrix(matrix, target))

	matrix1 := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	target1 := 13
	println(searchMatrix(matrix1, target1))
}
