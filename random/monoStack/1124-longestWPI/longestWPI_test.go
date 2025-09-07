package _124_longestWPI

import "testing"

func TestLongestWPI(t *testing.T) {
	hours1 := []int{9, 9, 6, 0, 6, 6, 9}
	println(longestWPI(hours1))

	hours2 := []int{0, 6, 6, 9}
	println(longestWPI(hours2))

	hours3 := []int{6, 0, 6, 6, 9, 9, 10, 7, 9}
	println(longestWPI(hours3))

	hours4 := []int{6, 6, 6}
	println(longestWPI(hours4))

	hours5 := []int{10}
	println(longestWPI(hours5))

	hours6 := []int{12, 8, 7, 7, 9, 10, 8, 7, 9, 7, 8, 11}
	println(longestWPI(hours6))

	hours7 := []int{7}
	println(longestWPI(hours7))
}
