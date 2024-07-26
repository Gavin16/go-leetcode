package _99_calcEquation

import (
	"fmt"
	"testing"
)

func TestCalcEquation(t *testing.T) {

	equations := [][]string{{"a", "b"}, {"b", "c"}}
	values := []float64{2.0, 3.0}
	queries := [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}

	equation := calcEquation(equations, values, queries)
	fmt.Printf("%v\n", equation)
}
