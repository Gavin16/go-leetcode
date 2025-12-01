package _774_closestCost

import "testing"

func TestClosestCost(t *testing.T) {
	tests := []struct {
		name         string
		baseCosts    []int
		toppingCosts []int
		target       int
		want         int
	}{
		{"1", []int{2, 3}, []int{4, 5, 100}, 18, 17}, // TODO: add test cases
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := closestCost(tt.baseCosts, tt.toppingCosts, tt.target)
			if got != tt.want {
				t.Errorf("(%v) = %v,%v,%v,%v , want %v", tt, tt.baseCosts, tt.toppingCosts, tt.target, got, tt.want)
			}
		})
	}
}
