package _30_solve

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	board := [][]byte{{'X', 'O', 'X'}, {'O', 'X', 'O'}, {'X', 'O', 'X'}}
	solve(board)
	for _, v := range board {
		fmt.Printf("%c\n", v)
	}
}
