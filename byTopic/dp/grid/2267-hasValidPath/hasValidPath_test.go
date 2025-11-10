package _267_hasValidPath

import "testing"

func TestHasValidPath(t *testing.T) {
	tests := []struct {
		name  string
		input [][]byte
		want  bool
	}{
		// TODO: add test cases
		{"1", [][]byte{
			{'(', '(', ')', ')'},
			{'(', ')', '(', ')'},
			{'(', ')', ')', ')'},
			{'(', ')', ')', ')'},
			{')', '(', ')', ')'}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hasValidPath(tt.input)
			if got != tt.want {
				t.Errorf("HasValidPath(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
