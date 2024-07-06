package _1_maxArea

import (
	"fmt"
	"testing"
)

func TestMaxArea(t *testing.T) {

	height1 := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	area1 := maxArea(height1)
	fmt.Println(area1)

	height2 := []int{1, 1}
	area2 := maxArea(height2)
	fmt.Println(area2)

}
