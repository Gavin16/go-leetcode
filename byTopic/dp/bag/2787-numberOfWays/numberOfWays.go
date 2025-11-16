package _787_numberOfWays

// 2787. 将一个数字表示成幂的和的方案数
// 给你两个 正 整数 n 和 x 。
// 请你返回将 n 表示成一些 互不相同 正整数的 x 次幂之和的方案数。
// 换句话说，你需要返回互不相同整数 [n1, n2, ..., nk] 的集合数目，满足 n = n1x + n2x + ... + nkx 。
//
// 由于答案可能非常大，请你将它对 109 + 7 取余后返回。
//
// 比方说，n = 160 且 x = 3 ，一个表示 n 的方法是 n = 2^3 + 3^3 + 5^3 。
//
// 示例 1：
// 输入：n = 10, x = 2
// 输出：1
// 解释：我们可以将 n 表示为：n = 3^2 + 1^2 = 10 。
// 这是唯一将 10 表达成不同整数 2 次方之和的方案。
//
// 示例 2：
// 输入：n = 4, x = 1
// 输出：2
// 解释：我们可以将 n 按以下方案表示：
// - n = 4^1 = 4 。
// - n = 3^1 + 1^1 = 4 。
//
// 提示：
// 1 <= n <= 300
// 1 <= x <= 5

func numberOfWays(n int, x int) int {
	memo := make([][]int, n+1)
	for i := range memo {
		memo[i] = make([]int, n+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	mod := int(1e9 + 7)
	var dfs func(i, left int) int
	dfs = func(i, left int) (ans int) {
		if left == 0 {
			return 1
		}
		p := &memo[i][left]
		if *p != -1 {
			return *p
		}
		defer func() { *p = ans }()
		pw := pow(i, x)
		cnt := 0
		if i > 0 {
			if pw <= left {
				cnt = (cnt + dfs(i-1, left-pw) + dfs(i-1, left)) % mod
			} else {
				cnt = (cnt + dfs(i-1, left)) % mod
			}
		}
		return cnt
	}
	return dfs(n, n)
}

func numberOfWays1(n int, x int) int {
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	// f[i][l] 代表枚举自然到正整数i 时, 还剩 l的值需要通过求 x 次方来求和得到
	// f[i][l] 等于 f[i-1][l] + f[i-1][l-pow(i, x)]
	f[0][0] = 1
	mod := int(1e9 + 7)
	for i := range n {
		p := pow(i+1, x)
		for l := 0; l <= n; l++ {
			if p <= l {
				f[i+1][l] = (f[i][l] + f[i][l-p]) % mod
			} else {
				f[i+1][l] = f[i][l]
			}
		}
	}
	return f[n][n]
}

func pow(n, x int) int {
	ret := 1
	for range x {
		ret *= n
	}
	return ret
}
