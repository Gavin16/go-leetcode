package _18_change

import "testing"

func TestChange(t *testing.T) {
	amount, coins := 2, []int{3}
	println(change(amount, coins))
}
