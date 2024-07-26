package _10_findOrder

import (
	"fmt"
	"testing"
)

func TestFindOrder(t *testing.T) {

	numCourses, prerequisites := 2, [][]int{{1, 0}}
	order := findOrder(numCourses, prerequisites)
	fmt.Printf("%v\n", order)

	numCourses1, prerequisites1 := 4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}
	order1 := findOrder(numCourses1, prerequisites1)
	fmt.Printf("%v\n", order1)

}
