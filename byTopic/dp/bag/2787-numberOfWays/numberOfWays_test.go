package _787_numberOfWays

import "testing"

func TestNumberOfWays(t *testing.T) {
	tests := []struct {
		name string
		n    int
		x    int
		want int
	}{
		{"1", 10, 2, 1},
		{"2", 4, 1, 2},
		{"3", 3, 2, 0},
		{"4", 13, 2, 1},
		{"5", 100, 2, 3},
		{"6", 1, 2, 1},
		{"7", 2, 2, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numberOfWays1(tt.n, tt.x)
			if got != tt.want {
				t.Errorf("numberOfWays(%v, %v) = %v, want %v", tt.n, tt.x, got, tt.want)
			}
		})
	}
}
