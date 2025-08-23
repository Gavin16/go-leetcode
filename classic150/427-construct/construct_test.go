package _27_construct

import (
	"fmt"
	"testing"
)

func TestConstruct(t *testing.T) {

	grid1 := [][]int{{0, 1}, {1, 0}}
	node1 := construct(grid1)
	fmt.Println(node1)
}
