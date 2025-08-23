package _31_minFallingPathSum

import "testing"

func TestMinFallingPathSum(t *testing.T) {

	matrix := [][]int{{100, -42, -46, -41}, {31, 97, 10, -10}, {-58, -51, 82, 89}, {51, 81, 69, -51}}

	println(minFallingPathSum(matrix))

}
