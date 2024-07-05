package _4_longestComPrefix

import (
	"fmt"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {

	strs := []string{"flower", "flow", "flight"}
	prefix := longestCommonPrefix(strs)
	fmt.Println(prefix)
}
