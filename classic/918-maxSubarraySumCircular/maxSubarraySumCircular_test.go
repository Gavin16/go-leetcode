package _18_maxSubarraySumCircular

import "testing"

func TestMaxSubarraySumCircular(t *testing.T) {
	//nums := []int{1, -6, 4, -3, 6}
	//println(maxSubarraySumCircular1(nums))
	//
	//nums1 := []int{1, -2, 3, -2}
	//println(maxSubarraySumCircular1(nums1))

	nums2 := []int{-5, 3, 5}
	println(maxSubarraySumCircular1(nums2))

	nums3 := []int{3, -2, 2, -3}
	println(maxSubarraySumCircular1(nums3))

	nums4 := []int{-2, -1, -3, 1, -4, 0, -5}
	println(maxSubarraySumCircular1(nums4))

	nums5 := []int{5, -1, 5, -2}
	println(maxSubarraySumCircular1(nums5))
}
