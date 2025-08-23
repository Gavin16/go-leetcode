package _575_countRoutes

import (
	"math"
	"testing"
)

func TestCountRoutes(t *testing.T) {
	locations := []int{2, 3, 6, 8, 4}
	start, finish, fuel := 1, 3, 5
	println(countRoutes(locations, start, finish, fuel))

	println(math.MaxInt32)
}
