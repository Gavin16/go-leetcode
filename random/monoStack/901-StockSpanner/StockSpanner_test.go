package _01_StockSpanner

import "testing"

func TestStockSpanner_Next(t *testing.T) {
	spanner := Constructor()
	println(spanner.Next(29))
	println(spanner.Next(91))
	println(spanner.Next(62))
	println(spanner.Next(76))
	println(spanner.Next(51))

	slice := []int{1, 2, 3, 4}
	println(slice[2:len(slice)])
	//println(spanner.Next(100))
	//println(spanner.Next(80))
	//println(spanner.Next(60))
	//println(spanner.Next(70))
	//println(spanner.Next(60))
	//println(spanner.Next(75))
	//println(spanner.Next(85))
}
