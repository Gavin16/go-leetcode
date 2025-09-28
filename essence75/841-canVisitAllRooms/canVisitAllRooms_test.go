package _41_canVisitAllRooms

import "testing"

func TestCanVisitAllRooms(t *testing.T) {
	//rooms := [][]int{{1}, {2}, {3}, {}}
	//println(canVisitAllRooms(rooms))
	rooms1 := [][]int{{1, 3}, {3, 0, 1}, {2}, {0}}
	println(canVisitAllRooms(rooms1))
	
}
