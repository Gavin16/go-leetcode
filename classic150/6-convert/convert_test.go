package __convert

import (
	"fmt"
	"testing"
)

func TestConvert(t *testing.T) {

	str1 := "PAYPALISHIRING"
	numRows1 := 3
	str := convert(str1, numRows1)
	fmt.Println(str)

	str2 := "ASASASAS"
	numRows2 := 2
	str2 = convert(str2, numRows2)
	fmt.Println(str2)

	str3 := "PAYPALISHIRING"
	numRows3 := 4
	str3 = convert(str3, numRows3)
	fmt.Println(str3)
}
