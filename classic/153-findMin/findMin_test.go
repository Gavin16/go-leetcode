package _53_findMin

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindMin(t *testing.T) {

	tests := []struct {
		name   string
		nums   []int
		expect int
	}{
		{
			name:   "旋转3次",
			nums:   []int{3, 4, 5, 1, 2},
			expect: 1,
		},
		{
			name:   "旋转0次",
			nums:   []int{1, 2, 3, 4, 5},
			expect: 1,
		},
		{
			name:   "旋转5次",
			nums:   []int{3, 4, 5, 6, 7, 1},
			expect: 1,
		},
		{
			name:   "旋转1次",
			nums:   []int{7, 1, 2, 3, 4, 5, 6},
			expect: 1,
		},
		{
			name:   "旋转4次",
			nums:   []int{11, 13, 15, 17},
			expect: 11,
		},
		{
			name:   "旋转4次",
			nums:   []int{11, 13, 15, 17},
			expect: 11,
		},
		{
			name:   "边界值",
			nums:   []int{1},
			expect: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := findMin(tt.nums)
			assert.Equal(t, tt.expect, actual)
		})
	}

}
