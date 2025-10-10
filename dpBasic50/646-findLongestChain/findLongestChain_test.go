package _46_findLongestChain

import "testing"

func TestFindLongestChain(t *testing.T) {
	//pairs := [][]int{{1, 2}, {7, 8}, {4, 5}, {2, 7}}
	//pairs := [][]int{{1, 2}, {2, 3}, {3, 4}}
	pairs := [][]int{{1, 2}, {2, 4}, {3, 5}, {4, 5}, {6, 7}}
	println(findLongestChain(pairs))
}
