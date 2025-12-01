package _405_longestDiverseString

import "testing"

func TestLongestDiverseString(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		c    int
		want string
	}{
		{"1", 7, 1, 1, "aabaacaa"}, // TODO: add test cases
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestDiverseString(tt.a, tt.b, tt.c)
			if got != tt.want {
				t.Errorf("(%v, %v, %v) = %v, want %v", tt.a, tt.b, tt.c, got, tt.want)
			}
		})
	}
}
