package _964_longestObstacleCourse

import (
	"fmt"
	"testing"
)

func TestLongestObstacleCourseAtEachPosition(t *testing.T) {
	obstacles := []int{5, 1, 5, 5, 1, 3, 4, 5, 1, 4}
	fmt.Println(longestObstacleCourseAtEachPosition(obstacles))
}
