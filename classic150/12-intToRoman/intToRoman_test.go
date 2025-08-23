package _2_intToRoman

import (
	"fmt"
	"testing"
)

func TestInToRoman(t *testing.T) {

	num := 3999
	roman := intToRoman1(num)
	fmt.Println(roman)

	num1 := 3749
	roman1 := intToRoman1(num1)
	fmt.Println(roman1)

	num2 := 1994
	roman2 := intToRoman1(num2)
	fmt.Println(roman2)

	num3 := 58
	roman3 := intToRoman(num3)
	fmt.Println(roman3)
}
