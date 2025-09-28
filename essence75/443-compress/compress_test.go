package _43_compress

import (
	"fmt"
	"testing"
)

func TestCompress(t *testing.T) {
	chars := []byte{'a', 'a'}

	println(compress(chars))
	for _, ch := range chars {
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	chars1 := []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}
	println(compress(chars1))
	for _, ch := range chars1 {
		fmt.Printf("%c ", ch)
	}
	fmt.Println()
}
