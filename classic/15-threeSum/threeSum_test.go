package _5_threeSum

import (
	"fmt"
	"testing"
)

func TestThreeSum(t *testing.T) {

	nums := []int{-1, 0, 1, 2, -1, -4}
	matches := threeSum(nums)
	for _, match := range matches {
		fmt.Printf("%v\n", match)
	}

	nums1 := []int{0, 1, 1}
	matches1 := threeSum(nums1)
	for _, match1 := range matches1 {
		fmt.Printf("%v\n", match1)
	}

	nums2 := []int{0, 0, 0}
	matches2 := threeSum(nums2)
	for _, match2 := range matches2 {
		fmt.Printf("%v\n", match2)
	}

}
