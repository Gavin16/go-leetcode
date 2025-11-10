package _4_minPathSum

import "testing"

func TestMinPathSum(t *testing.T) {
	tests := []struct {
		name  string
		input [][]int
		want  int
	}{
		{}, // TODO: add test cases
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minPathSum(tt.input)
			if got != tt.want {
				t.Errorf("minPathSum(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
