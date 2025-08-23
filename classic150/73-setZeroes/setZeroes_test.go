package _3_setZeroes

import (
	"fmt"
	"testing"
)

func TestSetZeros(t *testing.T) {

	matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(matrix)
	for _, row := range matrix {
		fmt.Printf("%v\n", row)
	}

	matrix1 := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroes(matrix1)
	for _, row := range matrix1 {
		fmt.Printf("%v\n", row)
	}

	matrix2 := [][]int{{0}}
	setZeroes(matrix2)
	for _, row := range matrix2 {
		fmt.Printf("%v\n", row)
	}

	matrix3 := [][]int{{1}}
	setZeroes(matrix3)
	for _, row := range matrix3 {
		fmt.Printf("%v\n", row)
	}
}
