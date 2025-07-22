package _22_coinChange

import (
	"fmt"
	"testing"
)

func TestCoinChange(t *testing.T) {

	coins1 := []int{1, 2, 5}
	amount1 := 11
	fmt.Println(coinChange2(coins1, amount1))

	coins2 := []int{2}
	amount2 := 3
	fmt.Println(coinChange2(coins2, amount2))

	coins3 := []int{1, 2, 5, 10, 20, 50, 100}
	amount3 := 999
	fmt.Println(coinChange2(coins3, amount3))

	coins4 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	amount4 := 100
	fmt.Println(coinChange2(coins4, amount4))

	coins5 := []int{1, 2147483647}
	amount5 := 2
	fmt.Println(coinChange2(coins5, amount5))

	coins6 := []int{2, 2147483647}
	amount6 := 1
	fmt.Println(coinChange2(coins6, amount6))

	coins7 := []int{456, 117, 5, 145}
	amount7 := 1459
	fmt.Println(coinChange2(coins7, amount7))

	coins8 := []int{186, 419, 83, 408}
	amount8 := 6249
	fmt.Println(coinChange2(coins8, amount8))
}
