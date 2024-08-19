package _5_searchInsert

import "testing"

func TestSearchInsert(t *testing.T) {

	nums := []int{1, 3, 5, 6}
	println(searchInsert(nums, 5))
	println(searchInsert(nums, 2))
	println(searchInsert(nums, 7))

	nums1 := []int{1, 3, 5, 6, 8, 10, 13}
	println(searchInsert(nums1, 7))
	println(searchInsert(nums1, 9))

	nums2 := []int{}
	println(searchInsert(nums2, 1))
	println(searchInsert(nums2, 2))
}
