package _466_minReorder

import "testing"

func TestMinReorder(t *testing.T) {
	conns := [][]int{{4, 5}, {0, 1}, {1, 3}, {2, 3}, {4, 0}}
	n := 6
	println(minReorder(n, conns))
}
