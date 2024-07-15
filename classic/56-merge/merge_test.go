package _6_merge

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {

	//intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	//ans := merge(intervals)
	//for _, v := range ans {
	//	fmt.Println(v)
	//}
	//
	//intervals1 := [][]int{{1, 4}, {4, 5}}
	//ans1 := merge(intervals1)
	//for _, v := range ans1 {
	//	fmt.Println(v)
	//}
	//
	//intervals2 := [][]int{{1, 3}}
	//ans2 := merge(intervals2)
	//for _, v := range ans2 {
	//	fmt.Println(v)
	//}
	//
	//intervals3 := [][]int{{1, 4}, {2, 3}}
	//ans3 := merge(intervals3)
	//for _, v := range ans3 {
	//	fmt.Println(v)
	//}
	//
	intervals4 := [][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}}
	ans4 := merge(intervals4)
	for _, v := range ans4 {
		fmt.Println(v)
	}

}
