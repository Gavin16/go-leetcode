package _99_calcEquation

import "testing"

func TestCalcEquation(t *testing.T) {
	equations := [][]string{{"a", "b"}, {"b", "c"}}
	values := []float64{2.0, 3.0}
	queries := [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"X", "X"}}
	res := calcEquation(equations, values, queries)
	for _, val := range res {
		println(val)
	}

}
