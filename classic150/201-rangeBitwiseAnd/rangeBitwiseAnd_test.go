package _01_rangeBitwiseAnd

import "testing"

func TestRangeBitwiseAnd(t *testing.T) {

	left, right := 5, 7
	println(rangeBitwiseAnd1(left, right))

	left1, right1 := 1, 2147483647
	println(rangeBitwiseAnd1(left1, right1))

}
