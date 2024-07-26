package _27_ladderLength

import "testing"

func TestLadderLength(t *testing.T) {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	println(ladderLength(beginWord, endWord, wordList))

	beginWord1 := "hit"
	endWord1 := "cog"
	wordList1 := []string{"hot", "dot", "dog", "lot", "log"}
	println(ladderLength(beginWord1, endWord1, wordList1))
}
