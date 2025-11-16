package _049_lastStoneWeightII

import "testing"

func TestLastStoneWeightII(t *testing.T) {
	tests := []struct {
		name   string
		stones []int
		want   int
	}{
		{"1", []int{1}, 1},
		{"2", []int{2, 7, 4, 1, 8, 1}, 1},
		{"3", []int{31, 26, 33, 21, 40}, 5},
		{"4", []int{1, 2, 3, 4}, 0},
		{"5", []int{1, 2, 3, 4, 5}, 1},
		{"6", []int{2}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lastStoneWeightII(tt.stones)
			if got != tt.want {
				t.Errorf("lastStoneWeightII(%v) = %v, want %v", tt.stones, got, tt.want)
			}
		})
	}
}
