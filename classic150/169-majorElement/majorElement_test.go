package _69_majorElement

import (
	"fmt"
	"testing"
)

func TestMajorElement(t *testing.T) {

	nums := []int{3, 2, 3}
	element := majorityElement(nums)
	fmt.Println(element)

	nums1 := []int{1}
	element1 := majorityElement(nums1)
	fmt.Println(element1)

	nums2 := []int{2, 2, 1, 1, 1, 2, 2}
	element2 := majorityElement(nums2)
	fmt.Println(element2)
}
