package _54_maxEnvelopes

import (
	"math/rand"
	"testing"
)

func TestMaxEnvelopes(t *testing.T) {
	//envelopes := generateEnvelopes(100000)
	//println(maxEnvelopes(envelopes))
	// [[2,100],[3,200],[4,300],[5,500],[5,400],[5,250],[6,370],[6,360],[7,380]]
	//envelopes := [][]int{{2, 100}, {3, 200}, {4, 300}, {5, 500}, {5, 400}, {5, 250}, {6, 370}, {6, 360}, {7, 380}}
	//envelopes := [][]int{{1, 1}, {1, 1}, {1, 1}}
	//envelopes := [][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}
	envelopes := [][]int{{15, 8}, {2, 20}, {2, 14}, {4, 17}, {8, 19}, {8, 9}, {5, 7}, {11, 19},
		{8, 11}, {13, 11}, {2, 13}, {11, 19}, {8, 11}, {13, 11}, {2, 13},
		{11, 19}, {16, 1}, {18, 13}, {14, 17}, {18, 19}}
	println(maxEnvelopes1(envelopes))
}

func generateEnvelopes(n int) [][]int {
	envelopes := make([][]int, 0)
	for i := 0; i < n; i++ {
		s, t := rand.Intn(n), rand.Intn(n)
		envelopes = append(envelopes, []int{s, t})
	}
	return envelopes
}
