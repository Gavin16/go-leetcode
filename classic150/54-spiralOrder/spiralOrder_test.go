package _4_spiralOrder

import (
	"fmt"
	"testing"
)

func TestSpiralOrder(t *testing.T) {

	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	order := spiralOrder(matrix)
	fmt.Printf("%v\n", order)

	matrix1 := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	order1 := spiralOrder(matrix1)
	fmt.Printf("%v\n", order1)

	matrix2 := [][]int{{1}}
	order2 := spiralOrder(matrix2)
	fmt.Printf("%v\n", order2)
}
