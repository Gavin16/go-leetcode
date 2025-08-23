package _7_combine

import (
	"fmt"
	"testing"
)

func TestCombine(t *testing.T) {
	n, k := 4, 2
	ans := combine(n, k)
	for _, v := range ans {
		fmt.Printf("%v ", v)
	}
	fmt.Println()

	n1, k1 := 1, 1
	ans1 := combine(n1, k1)
	for _, v := range ans1 {
		fmt.Printf("%v ", v)
	}
	fmt.Println()

}
