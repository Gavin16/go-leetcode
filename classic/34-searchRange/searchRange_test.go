package _4_searchRange

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchRange(t *testing.T) {

	tests := []struct {
		name   string
		nums   []int
		target int
		expect []int
	}{
		{
			name:   "正常情况",
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 8,
			expect: []int{3, 4},
		},
		{
			name:   "正常情况",
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 5,
			expect: []int{0, 0},
		},
		{
			name:   "正常情况",
			nums:   []int{5, 7, 7, 7, 7, 8, 8, 10},
			target: 7,
			expect: []int{1, 4},
		},
		{
			name:   "不存在情况",
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 6,
			expect: []int{-1, -1},
		},
		{
			name:   "正常情况",
			nums:   []int{},
			target: 3,
			expect: []int{-1, -1},
		},
		{
			name:   "正常情况",
			nums:   []int{1},
			target: 1,
			expect: []int{0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := searchRange(tt.nums, tt.target)
			assert.Equal(t, tt.expect, actual)
		})
	}

}
