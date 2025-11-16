package _94_findTargetSumWays

import "testing"

func TestFindTargetSumWays(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{"1", []int{1, 1, 1, 1, 1}, 3, 5},
		{"2", []int{0}, 0, 2},
		{"3", []int{1}, 1, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findTargetSumWays(tt.nums, tt.target)
			if got != tt.want {
				t.Errorf("findTargetSumWays(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
