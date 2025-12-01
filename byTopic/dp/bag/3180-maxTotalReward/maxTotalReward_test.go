package _180_maxTotalReward

import "testing"

func TestMaxTotalReward(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{"1", []int{1, 6, 4, 3, 2}, 11},
		{"2", []int{1, 1, 3, 3}, 4},
		{"3", []int{1, 2, 3, 4, 5, 6, 7, 8}, 15},
		{"4", []int{6}, 6},
		{"5", []int{6, 13, 9, 19}, 34},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxTotalReward2(tt.input)
			if got != tt.want {
				t.Errorf("maxTotalReward(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
