package _8_merge

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	m1, n1 := 3, len(nums2)
	merge(nums1, m1, nums2, n1)
	fmt.Println(nums1)

	nums3 := []int{1}
	nums4 := []int{}
	m2, n2 := 1, len(nums4)
	merge(nums3, m2, nums4, n2)
	fmt.Println(nums3)

	nums5 := []int{3, 4, 5, 6, 0, 0, 0, 0}
	nums6 := []int{5, 6, 6, 7}
	m3, n3 := 4, len(nums6)
	merge(nums5, m3, nums6, n3)
	fmt.Println(nums5)
}
