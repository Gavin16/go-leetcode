package _28_summaryRanges

import (
	"fmt"
	"testing"
)

func TestSummaryRanges(t *testing.T) {

	nums := []int{0, 1, 2, 4, 5, 7}
	ans := summaryRanges(nums)
	for _, v := range ans {
		fmt.Println(v)
	}

	nums1 := []int{0, 2, 3, 4, 6, 8, 9}
	ans1 := summaryRanges(nums1)
	for _, v := range ans1 {
		fmt.Println(v)
	}
}
