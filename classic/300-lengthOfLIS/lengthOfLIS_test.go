package _00_lengthOfLIS

import "testing"

func TestLengthOfLIS(t *testing.T) {

	nums1 := []int{10, 9, 2, 5, 3, 7, 101, 18}
	println(lengthOfLIS2(nums1))

	nums2 := []int{0}
	println(lengthOfLIS2(nums2))

	num3 := []int{0, 1, 0, 3, 2, 3}
	println(lengthOfLIS2(num3))

	num4 := []int{1, 3, 6, 7, 9, 4, 10, 5, 6}
	println(lengthOfLIS2(num4))

	res := lengthOfLIS2([]int{10, 9, 2, 5, 3, 7, 101, 18})
	println(res)

	res1 := lengthOfLIS2([]int{0, 1, 0, 3, 2, 3})
	println(res1)

}
