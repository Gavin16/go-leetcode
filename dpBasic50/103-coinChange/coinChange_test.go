package _03_coinChange

import "testing"

func TestCoinChange(t *testing.T) {
	coins := []int{1, 2, 5, 10, 20}
	amount := 99
	println(coinChange(coins, amount))
}
