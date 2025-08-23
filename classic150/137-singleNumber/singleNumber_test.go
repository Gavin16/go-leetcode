package _37_singleNumber

import "testing"

func TestSingleNumber(t *testing.T) {

	nums1 := []int{2, 2, 3, 2}
	nums2 := []int{0, 1, 0, 1, 0, 1, 99}
	nums3 := []int{-2, -2, 1, 1, 4, 1, 4, 4, -4, -2}

	println(singleNumber(nums1))
	println(singleNumber(nums2))
	println(singleNumber(nums3))

}
