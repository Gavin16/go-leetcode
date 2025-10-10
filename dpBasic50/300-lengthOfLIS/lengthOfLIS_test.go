package _00_lengthOfLIS

import (
	"math/rand"
	"testing"
)

func TestLengthOfLIS(t *testing.T) {
	nums := buildRandomNums(20)
	println(lengthOfLIS(nums))
}

func buildRandomNums(n int) []int {
	nums := make([]int, 0)
	for i := 0; i < n; i++ {
		num := rand.Intn(1000)
		nums = append(nums, num)
	}
	return nums
}
