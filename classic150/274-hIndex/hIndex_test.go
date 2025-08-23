package _74_hIndex

import (
	"fmt"
	"testing"
)

func TestHIndex(t *testing.T) {

	nums := []int{3, 0, 6, 1, 5}
	index := hIndex(nums)
	fmt.Println(index)

	nums1 := []int{1, 3, 1}
	index1 := hIndex(nums1)
	fmt.Println(index1)
}
