package _7_fullJustify

import (
	"fmt"
	"testing"
)

func TestFullJustify(t *testing.T) {

	words1 := []string{"This", "is", "an", "example", "of", "text", "justification."}
	maxLen1 := 16
	res1 := fullJustify(words1, maxLen1)
	for _, v := range res1 {
		fmt.Println(v)
	}

	words2 := []string{"What", "must", "be", "acknowledgment", "shall", "be"}
	maxLen2 := 16
	res2 := fullJustify(words2, maxLen2)
	for _, v := range res2 {
		fmt.Println(v)
	}

	words3 := []string{"Science", "is", "what", "we", "understand", "well", "enough", "to",
		"explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}
	maxLen3 := 20
	res3 := fullJustify(words3, maxLen3)
	for _, v := range res3 {
		fmt.Println(v)
	}

	words4 := []string{"aaa"}
	maxLen4 := 11
	res4 := fullJustify(words4, maxLen4)
	for _, v := range res4 {
		fmt.Println(v)
	}
}
