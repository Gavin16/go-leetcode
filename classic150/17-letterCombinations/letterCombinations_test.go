package _7_letterCombinations

import (
	"fmt"
	"testing"
)

func TestLetterCombinations(t *testing.T) {

	digits := "23"
	result := letterCombinations(digits)
	for _, combination := range result {
		fmt.Printf("%v ", combination)
	}
	fmt.Println()

	digits1 := ""
	result1 := letterCombinations(digits1)
	for _, combination := range result1 {
		fmt.Printf("%v ", combination)
	}
	fmt.Println()

	digits2 := "2"
	result2 := letterCombinations(digits2)
	for _, combination := range result2 {
		fmt.Printf("%v ", combination)
	}
	fmt.Println()

}
