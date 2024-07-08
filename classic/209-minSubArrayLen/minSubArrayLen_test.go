package _09_minSubArrayLen

import (
	"fmt"
	"testing"
)

func TestMinSubArrayLen(t *testing.T) {

	nums, target := []int{2, 3, 1, 2, 4, 3}, 7
	arrayLen := minSubArrayLen(target, nums)
	fmt.Println(arrayLen)

	nums1, target1 := []int{1, 4, 4}, 4
	arrayLen1 := minSubArrayLen(target1, nums1)
	fmt.Println(arrayLen1)

	nums2, target2 := []int{1, 1, 1, 1, 1, 1, 1, 1}, 11
	arrayLen2 := minSubArrayLen(target2, nums2)
	fmt.Println(arrayLen2)

	nums3, target3 := []int{1, 2, 3, 4, 5}, 11
	arrayLen3 := minSubArrayLen(target3, nums3)
	fmt.Println(arrayLen3)

}
