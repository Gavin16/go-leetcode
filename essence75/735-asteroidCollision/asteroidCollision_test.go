package _35_asteroidCollision

import (
	"fmt"
	"testing"
)

func TestAsteroidCollision(t *testing.T) {
	asteroids := []int{8, -8}
	ans := asteroidCollision(asteroids)
	for _, n := range ans {
		fmt.Printf("%d ", n)
	}
	fmt.Println()

	asteroids1 := []int{-2, -1, 1, 4, -3, -5}
	ans1 := asteroidCollision(asteroids1)
	for _, n := range ans1 {
		fmt.Printf("%d ", n)
	}
	fmt.Println()
}
