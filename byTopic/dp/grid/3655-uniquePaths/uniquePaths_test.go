package _655_uniquePaths

import "testing"

func TestUniquePaths(t *testing.T) {
	tests := []struct {
		name  string
		input [][]int
		want  int
	}{
		{"1", [][]int{{0, 1, 0}, {0, 0, 1}, {1, 0, 0}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := uniquePaths(tt.input)
			if got != tt.want {
				t.Errorf("UniquePaths(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
