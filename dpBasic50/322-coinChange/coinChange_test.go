package _22_coinChange

import "testing"

func TestCoinChange(t *testing.T) {
	coins := []int{2}
	amount := 3
	println(coinChange(coins, amount))
}
