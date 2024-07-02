package rmDuplicates2

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates2(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}
	res := removeDuplicates2(nums)
	fmt.Println(nums[:res])

	nums1 := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	res1 := removeDuplicates2(nums1)
	fmt.Println(nums1[:res1])

	nums2 := []int{1}
	res2 := removeDuplicates2(nums2)
	fmt.Println(nums2[:res2])

	nums3 := []int{3, 3, 3, 3, 3, 3, 3}
	res3 := removeDuplicates2(nums3)
	fmt.Println(nums3[:res3])
}
