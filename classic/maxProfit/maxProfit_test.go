package maxProfit

import "testing"

func TestMaxProfit(t *testing.T) {
	nums := []int{7, 1, 5, 3, 6, 4}
	expect := 5
	profit := maxProfit(nums)
	if expect != profit {
		t.Errorf("expect %d, got %d", expect, profit)
	}

	nums1 := []int{7, 6, 4, 3, 1}
	expect1 := 0
	profit1 := maxProfit(nums1)
	if expect1 != profit1 {
		t.Errorf("expect %d, got %d", expect1, profit1)
	}
}
