package _05_canPlaceFlowers

import "testing"

func TestCanPlaceFlowers(t *testing.T) {
	flowerbed, n := []int{1, 0, 0, 1}, 1
	println(canPlaceFlowers(flowerbed, n))
}
