package _16_canPartition

import (
	"testing"
)

func TestCanPartition(t *testing.T) {
	//nums := buildRandomNums(200)
	tests := []struct {
		name  string
		input []int
		want  bool
	}{
		//{"1", nums, true},
		{"2", []int{2, 3, 4, 5, 6, 7, 8, 9, 10}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := canPartition(tt.input)
			if got != tt.want {
				t.Errorf("canPartition(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func buildRandomNums(n int) []int {
	res := make([]int, 0)
	for i := 1; i <= n; i++ {
		res = append(res, 3)
	}
	return res
}
