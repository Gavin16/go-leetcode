package _16_removeDuplicateLetters

import "testing"

func TestRemoveDuplicateLetters(t *testing.T) {
	s := "cbacdcbc"
	//println(removeDuplicateLetters(s))

	s1 := "bcabc"
	//println(removeDuplicateLetters(s))

	s2 := "abacb"
	println(removeDuplicateLetters(s))
	println(removeDuplicateLetters(s1))
	println(removeDuplicateLetters(s2))

	//hash := make(map[byte]int)
	//hash[12] = 1
	//println(hash[12])
}
