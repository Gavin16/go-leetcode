package _6_rmDuplicates

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {

	nums := []int{1, 1, 2}
	r := removeDuplicates(nums)
	fmt.Println(nums[:r])

	nums1 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	r1 := removeDuplicates(nums1)
	fmt.Println(nums1[:r1])
}
