package _980_findDifferentBinaryString

import (
	"fmt"
	"testing"
)

func TestFindDifferentBinaryString(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  string
	}{
		{"1", []string{
			"111", "011", "001"}, ""}, // TODO: add test cases
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findDifferentBinaryString(tt.input)
			fmt.Println(got)
			set := make(map[string]bool)
			for _, str := range tt.input {
				set[str] = true
			}
			if len(got) != len(tt.input) || set[got] {
				t.Errorf("findDifferentBinaryString(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
	path := make([]int, 2, 4)
	path[0], path[1] = 1, 0
	newPath0 := append(path, 0)
	fmt.Println(newPath0)
	newPath1 := append(path, 1)
	fmt.Println(newPath0)
	fmt.Println(newPath1)

	val := 102
	fmt.Printf("%0*b", 16, val)
}
