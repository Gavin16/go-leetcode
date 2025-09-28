package _462_totalCost

import "testing"

func TestTotalCost(t *testing.T) {
	cost, k := []int{17, 12, 10, 2, 7, 2, 11, 20, 8}, 3
	candidates := 4
	println(totalCost(cost, k, candidates))

	cost1, k1 := []int{1, 3, 4}, 2
	candidates1 := 3
	println(totalCost(cost1, k1, candidates1))

	cost2, k2 := []int{28, 35, 21, 13, 21, 72, 35, 52, 74, 92, 25, 65, 77, 1, 73, 32, 43, 68, 8, 100, 84,
		80, 14, 88, 42, 53, 98, 69, 64, 40, 60, 23, 99, 83, 5, 21, 76, 34}, 32
	candidates2 := 12
	println(totalCost(cost2, k2, candidates2))
}
