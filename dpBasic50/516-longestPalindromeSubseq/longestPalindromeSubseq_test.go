package _16_longestPalindromeSubseq

import "testing"

func TestLongestPalindromeSubseq(t *testing.T) {
	s1 := "qsp"
	println(longestPalindromeSubseq(s1))
	s2 := "a"
	println(longestPalindromeSubseq(s2))
	s3 := "snfkaosirovjs"
	println(longestPalindromeSubseq(s3))
}
