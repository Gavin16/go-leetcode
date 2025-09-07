package _0_myPow

// 50. Pow(x, n)
// 实现 pow(x, n) ，即计算 x 的整数 n 次幂函数（即，xn ）。
//
// 示例 1：
//
// 输入：x = 2.00000, n = 10
// 输出：1024.00000
// 示例 2：
//
// 输入：x = 2.10000, n = 3
// 输出：9.26100
// 示例 3：
//
// 输入：x = 2.00000, n = -2
// 输出：0.25000
// 解释：2-2 = 1/22 = 1/4 = 0.25
//
// 提示：
//
// -100.0 < x < 100.0
// -231 <= n <= 231-1
// n 是一个整数
// 要么 x 不为零，要么 n > 0 。
// -104 <= xn <= 104

func myPow(x float64, n int) float64 {
	if n > 0 {
		return quickMul(x, n)
	} else {
		return 1 / quickMul(x, n)
	}
}

func quickMul(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	half := quickMul(x, n/2)
	if n%2 == 0 {
		return half * half
	}
	return half * half * x
}
