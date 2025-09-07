package _9_mySqrt

import (
	"math"
	"testing"
)

func TestMySqrt(t *testing.T) {
	//println(mySqrt(0))
	//println(mySqrt(1))
	println(mySqrt(2))
	println(mySqrt(4))
	println(mySqrt(8))
	println(mySqrt(9))
	println(mySqrt(int(math.Pow(2, 31) - 1)))
}
