package _606_maximumCostSubstring

import "testing"

func TestMaximumCostSubstring(t *testing.T) {
	s := "avsbdf"
	chars := "abcd"
	vals := []int{2, -1, 1, -100}
	println(maximumCostSubstring(s, chars, vals))
}
