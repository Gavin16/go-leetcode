package _41_cherryPickup

import "testing"

func TestCherryPickup(t *testing.T) {
	tests := []struct {
		name  string
		input [][]int
		want  int
	}{
		{"1", [][]int{{0, 1, -1}, {1, 0, -1}, {1, 1, 1}}, 5},
		{"2", [][]int{
			{0, 1, 1, 0, 0},
			{1, 1, 1, 1, 0},
			{-1, 1, 1, 1, -1},
			{0, 1, 1, 1, 0},
			{1, 0, -1, 0, 0}}, 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cherryPickup(tt.input)
			if got != tt.want {
				t.Errorf("CherryPickup(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
