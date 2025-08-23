package _7_insert

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {

	intervals1 := [][]int{{1, 4}}
	newInterval1 := []int{2, 3}
	ans1 := insert(intervals1, newInterval1)
	for _, v := range ans1 {
		fmt.Printf("%v\n", v)
	}

	intervals2 := [][]int{{1, 3}, {4, 6}}
	newInterval2 := []int{2, 4}
	ans2 := insert(intervals2, newInterval2)
	for _, v := range ans2 {
		fmt.Printf("%v\n", v)
	}
}
