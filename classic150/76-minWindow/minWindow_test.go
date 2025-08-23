package _6_minWindow

import "testing"

func TestMinWindow(t *testing.T) {

	s1, t1 := "ADOBECODEBANC", "ABC"
	println(minWindow(s1, t1))

	s2, t2 := "a", "a"
	println(minWindow(s2, t2))

	s3, t3 := "a", "aa"
	println(minWindow(s3, t3))
}
