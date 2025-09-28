package _38_productExceptSelf

import "testing"

func TestProductExceptSelf(t *testing.T) {
	nums1 := []int{1, 2, 3, 4}
	for _, num := range productExceptSelf(nums1) {
		println(num)
	}

	nums2 := []int{-1, 1, 0, -3, 3}
	for _, num := range productExceptSelf(nums2) {
		println(num)
	}
}
