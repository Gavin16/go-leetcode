package _62_findPeakElement

import "testing"
import "github.com/stretchr/testify/assert"

func TestFindPeakElement(t *testing.T) {

	tests := []struct {
		name   string
		nums   []int
		expect int
	}{
		{
			name:   "正常情况",
			nums:   []int{1, 2, 3, 4, 5},
			expect: 4,
		},
		{
			name:   "正常情况",
			nums:   []int{1, 2, 3, 1},
			expect: 2,
		},
		{
			name:   "正常情况",
			nums:   []int{1, 2, 1, 3, 5, 6, 4},
			expect: 5,
		},
		{
			name:   "正常情况",
			nums:   []int{5, 4, 3, 2, 1},
			expect: 0,
		},
		{
			name:   "峰值在中间",
			nums:   []int{1, 2, 3, 2, 1},
			expect: 2,
		},
		{
			name:   "峰值在开头",
			nums:   []int{3, 2, 1},
			expect: 0,
		},
		{
			name:   "峰值在结尾",
			nums:   []int{1, 2, 3},
			expect: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := findPeakElement(tt.nums)
			assert.Equal(t, tt.expect, actual)
		})
	}

}
