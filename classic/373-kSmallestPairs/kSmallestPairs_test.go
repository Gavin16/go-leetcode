package _73_kSmallestPairs

import "testing"

func TestKSmallestPairs(t *testing.T) {

	nums1 := []int{1, 7, 11}
	nums2 := []int{2, 4, 6}
	k := 3

	res1 := kSmallestPairs(nums1, nums2, k)
	for _, arr := range res1 {
		println("[", arr[0], ",", arr[1], "]")
	}

	nums12 := []int{1, 1, 2}
	nums22 := []int{1, 2, 3}
	k2 := 2
	res2 := kSmallestPairs(nums12, nums22, k2)
	for _, arr := range res2 {
		println("[", arr[0], ",", arr[1], "]")
	}
}
