package LCP47_securityCheck

import "testing"

func TestSecurityCheck(t *testing.T) {
	tests := []struct {
		name       string
		capacities []int
		k          int
		want       int
	}{
		{"1", []int{2, 2, 3}, 2, 2},
		{"2", []int{3, 3}, 3, 0},
		{"3", []int{3}, 0, 1},
		{"4", []int{3, 2, 5, 2, 7}, 6, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := securityCheck(tt.capacities, tt.k)
			if got != tt.want {
				t.Errorf("(%v,%v) = %v, want %v", tt.capacities, tt.k, got, tt.want)
			}
		})
	}
}
