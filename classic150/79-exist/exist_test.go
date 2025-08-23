package _9_exist

import "testing"

func TestExist(t *testing.T) {

	//board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	//word := "ABCCED"
	//println(exist(board, word))

	board1 := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	word1 := "ABCB"
	println(exist(board1, word1))

}
