package _3_uniquePathsWithObstacles

import "testing"

func TestUniquePathsWithObstacles(t *testing.T) {
	grids := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	println(uniquePathsWithObstacles(grids))

	grids1 := [][]int{{1}}
	println(uniquePathsWithObstacles(grids1))

	grids2 := [][]int{{0, 1}, {1, 0}}
	println(uniquePathsWithObstacles(grids2))
}
