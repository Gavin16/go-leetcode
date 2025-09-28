package _493_longestSubarray

import "testing"

func TestLongestSubarray(t *testing.T) {
	nums := []int{1, 1, 0, 1}
	println(longestSubarray(nums))

	nums1 := []int{0, 1, 1, 1, 0, 1, 1, 0, 1}
	println(longestSubarray(nums1))

	nums2 := []int{1, 1, 1, 1}
	println(longestSubarray(nums2))

	nums3 := []int{1}
	println(longestSubarray(nums3))

	nums4 := []int{0, 1}
	println(longestSubarray(nums4))

	nums5 := []int{0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 1, 1}
	println(longestSubarray(nums5))

	nums6 := []int{1, 1, 0, 0, 1, 1, 1, 0, 1}
	println(longestSubarray(nums6))
}
