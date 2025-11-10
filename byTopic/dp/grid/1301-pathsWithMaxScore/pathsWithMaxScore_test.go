package _301_pathsWithMaxScore

import (
	"fmt"
	"testing"
)

func TestPathsWithMaxScore(t *testing.T) {
	//board := []string{"E23", "2X2", "12S"}
	board := []string{"E12", "1X1", "21S"}
	fmt.Println(pathsWithMaxScore(board))
}
