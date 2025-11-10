package _684_maxMoves

import (
	"fmt"
	"testing"
)

func TestMaxMoves(t *testing.T) {
	grid := [][]int{
		{11, 4, 10, 5},
		{5, 9, 9, 11},
		//{3, 8, 7, 1},
		//{10, 7, 6, 1},
	}
	fmt.Println(maxMoves1(grid))
}
