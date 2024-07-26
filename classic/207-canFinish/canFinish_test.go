package _07_canFinish

import (
	"fmt"
	"testing"
)

func TestCanFinish(t *testing.T) {

	numCourses := 2
	prerequisites := [][]int{{1, 0}}
	println(canFinish(numCourses, prerequisites))

	numCourses1 := 2
	prerequisites1 := [][]int{{1, 0}, {0, 1}}
	println(canFinish(numCourses1, prerequisites1))

	numCourses2 := 4
	prerequisites2 := [][]int{{1, 0}, {2, 0}, {3, 2}, {2, 1}}
	println(canFinish(numCourses2, prerequisites2))

	order, err := canFinish2(numCourses2, prerequisites2)
	if err == nil {
		fmt.Printf("%v\n", order)
	}

}
