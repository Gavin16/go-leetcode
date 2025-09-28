package _01_StockSpanner

import "testing"

func TestStockSpanner(t *testing.T) {
	spanner := Constructor()
	println(spanner.Next(8))
	println(spanner.Next(7))
	println(spanner.Next(5))
	println(spanner.Next(5))
	println(spanner.Next(4))
	println(spanner.Next(3))
}
