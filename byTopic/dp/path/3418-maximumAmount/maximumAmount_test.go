package _418_maximumAmount

import (
	"fmt"
	"testing"
)

func TestMaximumAmount(t *testing.T) {
	//coins := [][]int{
	//	{0, 1, 4},
	//	{1, 3, -3},
	//	{2, -3, 4},
	//}

	coins := [][]int{
		{-7, 12, 12, 13},
		{-6, 19, 19, -6},
		{9, -2, -10, 16},
		{-4, 14, -10, -9},
	}

	fmt.Println(maximumAmount(coins))
}
