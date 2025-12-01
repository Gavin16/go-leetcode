package _585_waysToReachTarget

import "testing"

func TestWaysToReachTarget(t *testing.T) {
	tests := []struct {
		name   string
		input  [][]int
		target int
		want   int
	}{
		{"1", [][]int{{6, 1}, {3, 2}, {2, 3}}, 6, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := waysToReachTarget(tt.target, tt.input)
			if got != tt.want {
				t.Errorf("(%v, %v) = %v, want %v", tt.target, tt.input, got, tt.want)
			}
		})
	}
}
