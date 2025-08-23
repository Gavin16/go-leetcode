package _34_completeCircuit

import (
	"fmt"
	"testing"
)

func TestCanCompleteCircuit(t *testing.T) {
	gas1 := []int{1, 2, 3, 4, 5}
	cost1 := []int{3, 4, 5, 1, 2}

	pos1 := canCompleteCircuit1(gas1, cost1)
	fmt.Println(pos1)

	gas2 := []int{2, 3, 4}
	cost2 := []int{3, 4, 3}
	pos2 := canCompleteCircuit1(gas2, cost2)
	fmt.Println(pos2)

	gas3 := []int{3, 5, 2, 7, 6, 4, 8}
	cost3 := []int{4, 4, 3, 6, 5, 6, 7}
	pos3 := canCompleteCircuit1(gas3, cost3)
	fmt.Println(pos3)
}
