package _53_carFleet

import "testing"

func TestCarFleet(t *testing.T) {
	target1, position1, speed1 := 12, []int{10, 8, 0, 5, 3}, []int{2, 4, 1, 1, 3}

	println(carFleet(target1, position1, speed1))

}
