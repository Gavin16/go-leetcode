package main

import "fmt"

func testModifyRef(s []int) {
	if len(s) > 0 {
		s[len(s)-1] = 123
	}
}

func testSliceLen(s []int) []int {
	return append(s, 999)
}

func main() {
	//s := []int{1, 2, 3, 4, 5}
	//testModifyRef(s)
	//for _, n := range s {
	//	fmt.Println(n)
	//}
	s := make([]int, 1, 2)
	testSliceLen(s)
	fmt.Println(s)
}
