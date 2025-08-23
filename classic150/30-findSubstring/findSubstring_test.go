package _0_findSubstring

import (
	"fmt"
	"testing"
)

func TestFindSubString(t *testing.T) {

	//s, words := "barfoothefoobarman", []string{"foo", "bar"}
	//startPos := findSubstring2(s, words)
	//fmt.Printf("%v\n", startPos)
	//
	//s1, words1 := "wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}
	//startPos1 := findSubstring2(s1, words1)
	//fmt.Printf("%v\n", startPos1)
	//
	//s2, words2 := "barfoofoobarthefoobarman", []string{"bar", "foo", "the"}
	//startPos2 := findSubstring2(s2, words2)
	//fmt.Printf("%v\n", startPos2)

	s3, words3 := "lingmindraboofooowingdingbarrwingmonkeypoundcake", []string{"fooo", "barr", "wing", "ding", "wing"}
	startPos3 := findSubstring2(s3, words3)
	fmt.Printf("%v\n", startPos3)

}
