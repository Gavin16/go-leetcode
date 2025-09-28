package _94_orangesRotting

import "testing"

func TestOrangesRotting(t *testing.T) {
	grid := [][]int{{2, 1, 1}, {1, 1, 1}, {0, 1, 2}}
	println(orangesRotting(grid))
}
