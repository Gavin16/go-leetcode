package _9_combinationSum

import (
	"fmt"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	candidates := []int{2, 3, 6, 7}
	target := 7
	ans := combinationSum(candidates, target)
	for _, v := range ans {
		fmt.Printf("%v ", v)
	}
	fmt.Println()

	candidates1 := []int{2, 3, 5}
	target1 := 8
	ans1 := combinationSum(candidates1, target1)
	for _, v := range ans1 {
		fmt.Printf("%v ", v)
	}
	fmt.Println()

	candidates2 := []int{2}
	target2 := 1
	ans2 := combinationSum(candidates2, target2)
	for _, v := range ans2 {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
}
