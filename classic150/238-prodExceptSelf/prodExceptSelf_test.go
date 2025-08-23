package _38_prodExceptSelf

import (
	"fmt"
	"testing"
)

func TestProdExceptSelf(t *testing.T) {

	nums1 := []int{1, 2, 3, 4}
	self1 := productExceptSelf2(nums1)
	fmt.Println(self1)

	nums2 := []int{-1, 1, 0, -3, 3}
	self2 := productExceptSelf2(nums2)
	fmt.Println(self2)
}
