package _15_findKthLargest

import "testing"

func TestFindKthLargest(t *testing.T) {
	nums, k := []int{7, 6, 5, 4, 3, 2, 1}, 5
	println(findKthLargest(nums, k))
}
