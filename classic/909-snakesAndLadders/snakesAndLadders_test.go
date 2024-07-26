package _09_snakesAndLadders

import "testing"

func TestSnakesAndLadders(t *testing.T) {

	//board := [][]int{{-1, -1, -1, -1, -1, -1}, {-1, -1, -1, -1, -1, -1}, {-1, -1, -1, -1, -1, -1},
	//	{-1, 35, -1, -1, 13, -1}, {-1, -1, -1, -1, -1, -1}, {-1, 15, -1, -1, -1, -1}}
	//println(snakesAndLadders(board))

	//board1 := [][]int{{-1, -1, -1},
	//	{-1, 9, 8},
	//	{-1, 8, 9}}
	//println(snakesAndLadders(board1))

	board2 := [][]int{
		{1, 1, -1},
		{1, 1, 1},
		{-1, 1, 1}}
	println(snakesAndLadders(board2))
}
