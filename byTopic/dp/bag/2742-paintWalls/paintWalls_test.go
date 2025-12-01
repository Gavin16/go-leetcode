package _742_paintWalls

import "testing"

func TestPaintWalls(t *testing.T) {
	tests := []struct {
		name  string
		cost  []int
		times []int
		want  int
	}{
		{"1", []int{1, 2, 3, 2}, []int{1, 2, 3, 2}, 3},
		{"2", []int{2, 3, 4, 2}, []int{1, 1, 1, 1}, 4},
		{"3", []int{2, 5, 3, 7}, []int{1, 2, 3, 2}, 3},
		{"3", []int{3, 6, 5, 7}, []int{2, 3, 2, 2}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := paintWalls1(tt.cost, tt.times)
			if got != tt.want {
				t.Errorf("paintWalls(%v, %v) = %v, want %v", tt.cost, tt.times, got, tt.want)
			}
		})
	}
}
