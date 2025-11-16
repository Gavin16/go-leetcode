package _489_minZeroArray

import "testing"

func TestMinZeroArray(t *testing.T) {
	tests := []struct {
		name    string
		nums    []int
		queries [][]int
		want    int
	}{
		{"1", []int{2, 0, 2}, [][]int{{0, 2, 1}, {0, 2, 1}, {1, 1, 3}}, 2},
		{"2", []int{4, 3, 2, 1}, [][]int{{1, 3, 2}, {0, 2, 1}}, -1},
		{"3", []int{1, 2, 3, 2, 1}, [][]int{{0, 1, 1}, {1, 2, 1}, {2, 3, 2}, {3, 4, 1}, {4, 4, 1}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minZeroArray(tt.nums, tt.queries)
			if got != tt.want {
				t.Errorf("(%v, %v) = %v, want %v", tt.nums, tt.queries, got, tt.want)
			}
		})
	}
}
