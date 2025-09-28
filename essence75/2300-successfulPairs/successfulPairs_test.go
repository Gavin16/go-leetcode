package _300_successfulPairs

import "testing"

func TestSuccessfulPairs(t *testing.T) {
	//spells := []int{10, 2, 1, 3}
	//potions := []int{1, 2, 3}
	//success := int64(9)
	//println(successfulPairs(spells, potions, success))

	queries := []string{"bbb", "cc", "aabcc", "cccccddef"}
	words := []string{"a", "aa", "aaa", "aaaa"}
	frequency := numSmallerByFrequency(queries, words)
	for _, num := range frequency {
		println(num)
	}
}
