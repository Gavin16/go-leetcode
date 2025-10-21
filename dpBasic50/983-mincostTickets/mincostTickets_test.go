package _83_mincostTickets

import "testing"

func TestMincostTickets(t *testing.T) {
	days := []int{1, 4, 6, 7, 8, 20}
	costs := []int{2, 7, 15}

	println(mincostTickets(days, costs))
}
