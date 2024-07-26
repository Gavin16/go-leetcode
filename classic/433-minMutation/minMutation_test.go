package _33_minMutation

import "testing"

func TestMinMutation(t *testing.T) {
	//start := "AAAAACCC"
	//end := "AACCCCCC"
	//bank := []string{"AAAACCCC", "AAACCCCC", "AACCCCCC"}
	//println(minMutation(start, end, bank))

	start1 := "AAAAAAAT"
	end1 := "CCCCCCCC"
	bank1 := []string{"AAAAAAAC", "AAAAAAAA", "CCCCCCCC"}
	println(minMutation(start1, end1, bank1))
}
