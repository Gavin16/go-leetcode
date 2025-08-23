package _5_canJump2

import (
	"fmt"
	"testing"
)

func TestJump(t *testing.T) {

	nums := []int{2, 3, 1, 1, 4}
	res := jump(nums)
	fmt.Println(res)

	nums1 := []int{2, 3, 0, 1, 4}
	res1 := jump(nums1)
	fmt.Println(res1)

	nums2 := []int{1, 2, 1, 1, 1}
	res2 := jump(nums2)
	fmt.Println(res2)
}
