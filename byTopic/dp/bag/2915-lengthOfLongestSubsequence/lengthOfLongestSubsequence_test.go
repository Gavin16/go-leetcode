package _915_lengthOfLongestSubsequence

import "testing"

func TestLengthOfLongestSubsequence(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		target int
		want   int
	}{
		{"1", []int{1, 2, 3, 4, 5}, 9, 3},
		{"2", []int{4, 1, 3, 2, 1, 5}, 7, 4},
		{"3", []int{1, 1, 5, 4, 5}, 3, -1},
		{"4", []int{1}, 1, 1},
		{"5", []int{7, 8, 6, 5, 11, 6, 4, 5, 3}, 29, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lengthOfLongestSubsequence1(tt.input, tt.target)
			if got != tt.want {
				t.Errorf("(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
