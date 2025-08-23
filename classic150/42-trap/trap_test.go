package _2_trap

import (
	"fmt"
	"testing"
)

func TestTrap(t *testing.T) {

	heights := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	sum := trap(heights)
	fmt.Println(sum)

	heights1 := []int{4, 2, 0, 3, 2, 5}
	sum1 := trap(heights1)
	fmt.Println(sum1)

	heights2 := []int{4, 2, 3}
	sum2 := trap(heights2)
	fmt.Println(sum2)
}
