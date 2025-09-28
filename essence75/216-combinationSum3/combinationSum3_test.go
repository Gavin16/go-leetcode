package _16_combinationSum3

import (
	"fmt"
	"testing"
)

func TestCombinationSum3(t *testing.T) {
	k, n := 9, 45
	sum := combinationSum3(k, n)
	for _, comb := range sum {
		for _, num := range comb {
			fmt.Printf("%d ", num)
		}
		fmt.Println()
	}
}
