package _15_findKthLargest

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKthLargest(t *testing.T) {

	tests := []struct {
		name   string
		nums   []int
		k      int
		expect int
	}{
		{
			name:   "正常情况",
			nums:   []int{3, 2, 1, 5, 6, 4},
			k:      2,
			expect: 5,
		},
		{
			name:   "正常情况",
			nums:   []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			k:      4,
			expect: 4,
		},
		{
			name:   "边界情况",
			nums:   []int{3},
			k:      1,
			expect: 3,
		},
		{
			name:   "边界情况",
			nums:   []int{3, 3, 3, 3, 3},
			k:      1,
			expect: 3,
		},
		{
			name:   "边界情况",
			nums:   []int{3, 3, 3, 3, 3},
			k:      5,
			expect: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := findKthLargest(tt.nums, tt.k)
			assert.Equal(t, tt.expect, actual)
		})
	}
}
