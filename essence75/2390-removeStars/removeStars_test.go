package _390_removeStars

import "testing"

func TestRemoveStars(t *testing.T) {
	s := "****sstr"
	println(removeStars(s))

	s1 := "1234567****sstr"
	println(removeStars(s1))
}
