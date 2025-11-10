package _304_minPathCost

import "testing"

func TestMinPathCost(t *testing.T) {
	grid := [][]int{{5, 3}, {4, 0}, {2, 1}}
	movCost := [][]int{{9, 8}, {1, 5}, {10, 12}, {18, 6}, {2, 4}, {14, 3}}
	println(minPathCost(grid, movCost))
}
