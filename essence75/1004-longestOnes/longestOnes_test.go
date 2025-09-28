package _004_longestOnes

import "testing"

func TestLongestOnes(t *testing.T) {
	nums, k := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2
	println(longestOnes(nums, k))

	nums1, k1 := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 0
	println(longestOnes(nums1, k1))

	nums2, k2 := []int{1, 1, 1, 0, 1, 0, 1, 1, 1, 1, 0}, 1
	println(longestOnes(nums2, k2))

	nums3, k3 := []int{0, 0, 1, 0, 1, 0, 1, 0, 0, 1, 0}, 2
	println(longestOnes(nums3, k3))
}
