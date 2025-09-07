package _02_findMaximizedCapital

import "testing"

func TestFindMaximizedCapital(t *testing.T) {
	//k, w, profits, capital := 2, 0, []int{1, 2, 3}, []int{0, 1, 1}
	//println(findMaximizedCapital(k, w, profits, capital))

	k, w, profits, capital := 3, 0, []int{1, 2, 3}, []int{0, 1, 2}
	println(findMaximizedCapital(k, w, profits, capital))
}
