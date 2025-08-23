package _22_maxProfit2

import (
	"fmt"
	"testing"
)

func TestMaxProfit(t *testing.T) {

	//nums := []int{7, 1, 5, 3, 6, 4}
	//profit := 121-maxProfit(nums)
	//fmt.Println(profit)

	nums := []int{7, 1, 5, 3, 6, 4}
	profit := maxProfit1(nums)
	fmt.Println(profit)

	nums1 := []int{1, 2, 3, 4, 5}
	profit1 := maxProfit1(nums1)
	fmt.Println(profit1)

	nums2 := []int{7, 6, 4, 3, 1}
	profit2 := maxProfit1(nums2)
	fmt.Println(profit2)
}
