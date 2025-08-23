package _67_twoSum

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	numbers, target := []int{2, 7, 11, 15}, 9
	sum := twoSum(numbers, target)
	fmt.Printf("%v\n", sum)

	numbers1, target1 := []int{2, 3, 4}, 6
	sum1 := twoSum(numbers1, target1)
	fmt.Printf("%v\n", sum1)
}
