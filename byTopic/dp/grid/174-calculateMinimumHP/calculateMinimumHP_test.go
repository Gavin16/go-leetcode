package _74_calculateMinimumHP

import "testing"

func TestCalculateMinimumHP(t *testing.T) {
	tests := []struct {
		name  string
		input [][]int
		want  int
	}{
		// TODO: add test cases
		{"1", [][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}}, 7},
		{"2", [][]int{{1, 2, 3}, {1, 1, 1}, {2, 3, 2}}, 1},
		{"3", [][]int{{0}}, 1},
		{"4", [][]int{{-1, 1, 1}, {-1, -1, 1}, {1, -1, -1}}, 2},
		{"5", [][]int{{0, -7, 5, -7}, {-3, -4, 3, -4}, {-9, -21, -4, 2}}, 8},
		{"6", [][]int{{2, 3, -1}, {4, 3, -16}, {-1, -16, 2}}, 8},
		{"7", [][]int{{1, -3, -1, -10}, {-8, 8, -2, -5}, {-10, -10, -5, 1}}, 3},
		{"7", [][]int{{1, -3, 3}, {0, -2, 0}, {-3, -3, -3}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculateMinimumHP(tt.input)
			if got != tt.want {
				t.Errorf("CalculateMinimumHP(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
