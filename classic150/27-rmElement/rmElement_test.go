package _7_rmElement

import (
	"fmt"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	val := 3
	k := removeElement(nums, val)
	fmt.Printf("k=%v, nums=%v\n", k, nums[:k])

	nums1 := []int{0, 1, 2, 3, 0, 4, 2}
	val1 := 2
	k1 := removeElement2(nums1, val1)
	fmt.Printf("k1=%v, nums1=%v\n", k1, nums1[:k1])

	nums2 := []int{2, 2, 2, 2}
	val2 := 2
	k2 := removeElement2(nums2, val2)
	fmt.Printf("k2=%v, nums2=%v\n", k2, nums2[:k2])

	var nums3 []int
	val3 := 1
	k3 := removeElement2(nums3, val3)
	fmt.Printf("k3=%v, nums3=%v\n", k3, nums3[:k3])
}
