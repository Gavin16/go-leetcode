package _47_findCircleNum

import "testing"

func TestFindCircleNum(t *testing.T) {
	matrix := [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}
	println(findCircleNum(matrix))
}
