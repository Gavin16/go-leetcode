package _693_climbStairs

import "testing"

func TestClimbStairs(t *testing.T) {
	n, costs := 4, []int{1, 2, 3, 4}
	println(climbStairs(n, costs))
}
