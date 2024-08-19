package _2_generateParenthesis

import (
	"fmt"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	n := 2
	parenthesis := generateParenthesis(n)
	for _, v := range parenthesis {
		fmt.Printf("%v\n", v)
	}

	n1 := 3
	parenthesis1 := generateParenthesis(n1)
	for _, v := range parenthesis1 {
		fmt.Printf("%v\n", v)
	}

	n2 := 1
	parenthesis2 := generateParenthesis(n2)
	for _, v := range parenthesis2 {
		fmt.Printf("%v\n", v)
	}

}
