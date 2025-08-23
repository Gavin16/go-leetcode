package _98_rob

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRob(t *testing.T) {

	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"1", []int{1, 2, 3, 1}, 4},
		{"2", []int{2, 7, 9, 3, 1}, 12},
		{"3", []int{2, 1, 1, 2}, 4},
		{"4", []int{1, 2}, 2},
		{"4", []int{3}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := rob(tt.nums)
			assert.Equal(t, tt.want, actual)
		})
	}
}
