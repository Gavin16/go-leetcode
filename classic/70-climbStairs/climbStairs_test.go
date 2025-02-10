package _0_climbStairs

import "testing"

func TestClimbStairs(t *testing.T) {

	tests := []struct {
		name     string
		n        int
		expected int
	}{
		{
			name:     "case1",
			n:        1,
			expected: 1,
		},
		{
			name:     "case2",
			n:        2,
			expected: 2,
		},
		{
			name:     "case3",
			n:        3,
			expected: 3,
		},
		{
			name:     "case3",
			n:        4,
			expected: 5,
		},
		{
			name:     "case5",
			n:        5,
			expected: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := climbStairs(tt.n)
			if actual != tt.expected {
				t.Errorf("climbStairs() = %v, want %v", actual, tt.expected)
			}
		})
	}

}
