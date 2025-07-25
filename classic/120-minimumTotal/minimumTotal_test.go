package _20_minimumTotal

import "testing"

func TestMinimumTotalT(t *testing.T) {

	triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	res1 := minimumTotal(triangle)
	println(res1)

}
