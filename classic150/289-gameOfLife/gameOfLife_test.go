package _89_gameOfLife

import (
	"fmt"
	"testing"
)

func TestGameOfLife(t *testing.T) {

	board := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	gameOfLife(board)
	for _, row := range board {
		fmt.Printf("%v\n", row)
	}

	board1 := [][]int{{1, 1}, {1, 0}}
	gameOfLife(board1)
	for _, row := range board1 {
		fmt.Printf("%v\n", row)
	}

	board2 := [][]int{{0}}
	gameOfLife(board2)
	for _, row := range board2 {
		fmt.Printf("%v\n", row)
	}

	board3 := [][]int{{1}}
	gameOfLife(board3)
	for _, row := range board3 {
		fmt.Printf("%v\n", row)
	}
}
