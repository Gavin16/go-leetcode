package _3_maxSubArray

import (
	"fmt"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	nums1 := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	max1 := maxSubArray(nums1)
	fmt.Println(max1)

	nums2 := []int{5, 4, -1, 7, 8}
	max2 := maxSubArray(nums2)
	fmt.Println(max2)

	nums3 := []int{2, -5, 4, 5, 6, 7, 8, 9, 10}
	max3 := maxSubArray(nums3)
	fmt.Println(max3)

	nums4 := []int{5, -3, 5, 5, -3, 5}
	max4 := maxSubArray(nums4)
	fmt.Println(max4)
}
