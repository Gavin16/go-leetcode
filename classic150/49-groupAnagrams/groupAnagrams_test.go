package _9_groupAnagrams

import (
	"fmt"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	groups := groupAnagrams(strs)
	for _, strGroup := range groups {
		fmt.Printf("%v\n", strGroup)
	}

	strs1 := []string{"a"}
	groups1 := groupAnagrams(strs1)
	for _, strGroup := range groups1 {
		fmt.Printf("%v\n", strGroup)
	}

	strs2 := []string{""}
	groups2 := groupAnagrams(strs2)
	for _, strGroup := range groups2 {
		fmt.Printf("%v\n", strGroup)
	}
}
