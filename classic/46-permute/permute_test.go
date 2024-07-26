package _6_permute

import (
	"fmt"
	"testing"
)

func TestPermute(t *testing.T) {
	nums := []int{1, 2, 3}
	ans := permute(nums)
	for _, v := range ans {
		fmt.Printf("%v ", v)
	}
	fmt.Println()

	nums1 := []int{0, 1}
	ans1 := permute(nums1)
	for _, v := range ans1 {
		fmt.Printf("%v ", v)
	}
	fmt.Println()

	nums2 := []int{1}
	ans2 := permute(nums2)
	for _, v := range ans2 {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
}
