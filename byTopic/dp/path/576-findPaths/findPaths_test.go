package _76_findPaths

import "testing"

func TestFindPaths(t *testing.T) {
	//m = 2, n = 2, maxMove = 2, startRow = 0, startColumn = 0
	//m, n := 2, 2
	//maxMove := 2
	//startRow, startColumn := 0, 0
	//println(findPaths(m, n, maxMove, startRow, startColumn))

	// m = 1, n = 3, maxMove = 3, startRow = 0, startColumn = 1

	m, n, maxMove, startRow, startColumn := 2, 2, 2, 0, 0
	//
	//m, n, maxMove, startRow, startColumn := 1, 3, 3, 0, 1
	println(findPaths2(m, n, maxMove, startRow, startColumn))
}
