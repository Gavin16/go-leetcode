package _0_isValid

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	s := "()[]{}"
	println(isValid(s))

	s1 := "(()[]{}"
	println(isValid(s1))

	s2 := "([]{})()"
	println(isValid(s2))

}
