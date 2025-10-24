package backtrace

import (
	"fmt"
	"testing"
)

func TestNthSuperUglyNumber(t *testing.T) {
	//primes := GeneratePrimes(100)
	primes := []int{2, 7, 13, 19}
	fmt.Println(nthSuperUglyNumber(12, primes))
}

// 生成 [2, n] 内的所有质数
func GeneratePrimes(n int) []int {
	sieve := make([]bool, n+1)
	var primes []int
	for i := 2; i <= n; i++ {
		if !sieve[i] {
			primes = append(primes, i)
			for j := i * i; j <= n; j += i {
				sieve[j] = true
			}
		}
	}
	return primes
}
