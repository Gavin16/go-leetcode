package _35_candy

import (
	"fmt"
	"testing"
)

func TestCandy(t *testing.T) {

	ratings := []int{2, 3, 1, 2, 2, 4, 5, 5, 3}
	res := candy1(ratings)
	fmt.Println(res)

	ratings1 := []int{1, 0, 2}
	res1 := candy1(ratings1)
	fmt.Println(res1)

	ratings2 := []int{1, 2, 2}
	res2 := candy1(ratings2)
	fmt.Println(res2)

	ratings3 := []int{1, 2}
	res3 := candy1(ratings3)
	fmt.Println(res3)

	ratings4 := []int{1, 3, 4, 5, 2}
	res4 := candy1(ratings4)
	fmt.Println(res4)
}
