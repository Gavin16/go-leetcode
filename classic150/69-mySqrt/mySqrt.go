package _9_mySqrt

import "math"

func mySqrt(x int) int {
	// 二分法找出 n*n < x 且(n+1)*(n+1) > x 的n
	low, hi := 0, int(math.Pow(2, 16))

	for low < hi {
		mid := (low + hi) / 2
		if mid*mid == x {
			return mid
		} else if mid*mid < 0 || mid*mid > x {
			hi = mid
		} else {
			low = mid + 1
		}
	}
	return low - 1
}
