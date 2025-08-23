package _25_isPalindrome

import "testing"

func TestIsPalindrome(t *testing.T) {

	s1 := "A man, a plan, a canal: Panama"
	s2 := "race a car"
	s3 := " "

	s4 := "0P"

	println(isPalindrome(s1))
	println(isPalindrome(s2))
	println(isPalindrome(s3))
	println(isPalindrome(s4))
}
