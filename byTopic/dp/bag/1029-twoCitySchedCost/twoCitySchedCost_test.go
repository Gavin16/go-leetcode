package _029_twoCitySchedCost

import "testing"

func TestTwoCitySchedCost(t *testing.T) {
	tests := []struct {
		name  string
		input [][]int
		want  int
	}{
		{"1", [][]int{{259, 770}, {448, 54}, {926, 667}, {184, 139}, {840, 118}, {577, 469}}, 1859},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoCitySchedCost(tt.input)
			if got != tt.want {
				t.Errorf("(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
