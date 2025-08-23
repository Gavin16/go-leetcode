package _8_rotate

import (
	"fmt"
	"testing"
)

func TestRotate(t *testing.T) {

	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(matrix)
	for _, row := range matrix {
		fmt.Printf("%v\n", row)
	}

	matrix1 := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	rotate(matrix1)
	for _, row := range matrix1 {
		fmt.Printf("%v\n", row)
	}

	matrix2 := [][]int{{1, 2}, {3, 4}}
	rotate(matrix2)
	for _, row := range matrix2 {
		fmt.Printf("%v\n", row)
	}

	matrix3 := [][]int{{3}}
	rotate(matrix3)
	for _, row := range matrix3 {
		fmt.Printf("%v\n", row)
	}
}
