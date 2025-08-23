package _8_strStr

import "testing"

func TestStrStr(t *testing.T) {
	haystack, needle := "sadbutsad", "sad"
	pos := strStr(haystack, needle)
	println(pos)

	haystack1, needle1 := "leetcode", "et"
	pos1 := strStr(haystack1, needle1)
	println(pos1)

	haystack2, needle2 := "hell", "hello"
	pos2 := strStr(haystack2, needle2)
	println(pos2)

	haystack3, needle3 := "mississippi", "issip"
	pos3 := strStr(haystack3, needle3)
	println(pos3)
}
