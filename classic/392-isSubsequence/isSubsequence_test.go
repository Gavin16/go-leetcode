package _92_isSubsequence

import "testing"

func TestIsSubsequence(t *testing.T) {

	s1, t1 := "abc", "ahbgdc"
	println(isSubsequence(s1, t1))

	s2, t2 := "axc", "ahbgdc"
	println(isSubsequence(s2, t2))
}
