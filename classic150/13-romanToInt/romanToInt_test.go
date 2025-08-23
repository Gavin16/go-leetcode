package _3_romanToInt

import (
	"fmt"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	//s := "XI"
	//toInt := romanToInt(s)
	//fmt.Println(toInt)

	s1 := "III"
	s2 := "IV"
	s3 := "IX"
	s4 := "LVIII"
	s5 := "MCMXCIV"

	fmt.Println(romanToInt(s1))
	fmt.Println(romanToInt(s2))
	fmt.Println(romanToInt(s3))
	fmt.Println(romanToInt(s4))
	fmt.Println(romanToInt(s5))

}
