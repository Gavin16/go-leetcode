package _89_rotate

import (
	"testing"
)

func TestRotate(t *testing.T) {

	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate2(nums, k)
	expected := []int{5, 6, 7, 1, 2, 3, 4}
	for i, v := range expected {
		if nums[i] != v {
			t.Errorf("189-rotate failed at index %d expected %d, got %d", i, v, nums[i])
		}
	}

	nums1 := []int{-1, -100, 3, 99}
	k1 := 2
	rotate2(nums1, k1)
	expected1 := []int{3, 99, -1, -100}
	for i, v := range expected1 {
		if nums1[i] != v {
			t.Errorf("189-rotate failed at index %d expected %d, got %d", i, v, nums1[i])
		}
	}
}
