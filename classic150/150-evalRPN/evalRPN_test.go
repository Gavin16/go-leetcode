package _50_evalRPN

import "testing"

func TestEvalRPN(t *testing.T) {

	tokens := []string{"2", "1", "+", "3", "*"}
	println(evalRPN(tokens))

	tokens1 := []string{"4", "13", "5", "/", "+"}
	println(evalRPN(tokens1))

	tokens2 := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	println(evalRPN(tokens2))

	tokens3 := []string{"10"}
	println(evalRPN(tokens3))

}
