package _679_maxOperations

import "testing"

func TestMaxOperations(t *testing.T) {
	nums := []int{2, 5, 4, 4, 1, 3, 4, 4, 1, 4, 4, 1, 2, 1, 2, 2, 3, 2, 4, 2}
	k := 3
	println(maxOperations(nums, k))
}
