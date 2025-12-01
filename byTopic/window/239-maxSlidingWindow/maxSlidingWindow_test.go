package _39_maxSlidingWindow

import (
	"slices"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		k     int
		want  []int
	}{
		{"1", []int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []int{3, 3, 5, 5, 6, 7}},
		{"2", []int{1}, 1, []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxSlidingWindow(tt.input, tt.k)
			if !slices.Equal(got, tt.want) {
				t.Errorf("(%v) = %v, %v, want %v", tt.input, tt.k, got, tt.want)
			}
		})
	}
}
