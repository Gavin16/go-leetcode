package _3_search

import "testing"

func TestSearch(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		expect int
	}{
		{
			name:   "正常情况",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 0,
			expect: 4,
		},
		{
			name:   "不存在情况",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 3,
			expect: -1,
		},
		{
			name:   "不存在情况",
			nums:   []int{1},
			target: 0,
			expect: -1,
		},
		{
			name:   "正常情况",
			nums:   []int{1, 2, 3, 4, 5},
			target: 1,
			expect: 0,
		},
		{
			name:   "正常情况",
			nums:   []int{1, 2, 3, 4, 5},
			target: 5,
			expect: 4,
		},
		{
			name:   "正常情况",
			nums:   []int{3},
			target: 3,
			expect: 0,
		},
		{
			name:   "正常情况",
			nums:   []int{5, 1, 3},
			target: 5,
			expect: 0,
		},
		{
			name:   "正常情况",
			nums:   []int{4, 5, 6, 7, 8, 1, 2, 3},
			target: 8,
			expect: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := search(tt.nums, tt.target)
			if actual != tt.expect {
				t.Errorf("expect: %v, actual: %v", tt.expect, actual)
			}
		})
	}
}
