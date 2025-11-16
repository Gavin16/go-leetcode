package _685_subsequenceSumAfterCapping

import (
	"testing"
)

func TestSubsequenceSumAfterCapping(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		k     int
		want  []bool
	}{
		{"1", []int{4, 3, 2, 4}, 5, []bool{false, false, true, true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := subsequenceSumAfterCapping(tt.input, tt.k)
			for i, b := range got {
				if b != tt.want[i] {
					t.Errorf("(%v) = %v, want %v", tt.input, got, tt.want)
					break
				}
			}
		})
	}
}
