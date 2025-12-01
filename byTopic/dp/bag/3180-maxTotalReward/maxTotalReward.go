package _180_maxTotalReward

import (
	"math/big"
	"slices"
	"sort"
)

// 3180. 执行操作可获得的最大总奖励 I
// 给你一个整数数组 rewardValues，长度为 n，代表奖励的值。
//
// 最初，你的总奖励 x 为 0，所有下标都是 未标记 的。你可以执行以下操作 任意次 ：
//
// 从区间 [0, n - 1] 中选择一个 未标记 的下标 i。
// 如果 rewardValues[i] 大于 你当前的总奖励 x，则将 rewardValues[i] 加到 x 上（即 x = x + rewardValues[i]），并 标记 下标 i。
// 以整数形式返回执行最优操作能够获得的 最大 总奖励。
//
// 示例 1：
// 输入：rewardValues = [1,1,3,3]
// 输出：4
//
// 解释：
// 依次标记下标 0 和 2，总奖励为 4，这是可获得的最大值。
//
// 示例 2：
// 输入：rewardValues = [1,6,4,3,2]
// 输出：11
//
// 解释：
// 依次标记下标 0、2 和 1。总奖励为 11，这是可获得的最大值。
//
// 提示：
// 1 <= rewardValues.length <= 2000
// 1 <= rewardValues[i] <= 2000

func maxTotalReward(rewardValues []int) int {
	// 记f[i][j] 代表以i元素结尾,是否能够获得最大奖励j
	// 则 f[i][j] = f[i-1][j-v] || f[i-1][j]  (j >= v && j - v < v)
	// 记 rewardValues 的最大值为m, 则 j最大能取到 2*m-1,  边界情况就是 取m前刚好总奖励为 m-1
	// f[0][rewardValues[0]] = rewardValues[0]
	slices.Sort(rewardValues)
	rewardValues = slices.Compact(rewardValues)
	n := len(rewardValues)
	m := rewardValues[n-1]
	f := make([]bool, 2*m)
	f[0] = true
	for _, v := range rewardValues {
		for j := range f {
			if j >= v && j < 2*v {
				f[j] = f[j] || f[j-v]
			}
		}
	}
	res := 0
	for i := 0; i < len(f); i++ {
		if f[i] {
			res = i
		}
	}
	return res
}

// 使用bitset 进行优化
func maxTotalReward1(rewardValues []int) int {
	slices.Sort(rewardValues)
	rewardValues = slices.Compact(rewardValues)
	one := big.NewInt(1)
	f := big.NewInt(1)
	p := new(big.Int)

	// 对布尔值求或运算, 对应在bitset 中为对bit 求 or 运算
	// big.Int 求Int.Or 是
	for _, v := range rewardValues {
		mask := p.Sub(p.Lsh(one, uint(v)), one)
		f.Or(f, p.Lsh(p.And(f, mask), uint(v)))
	}
	return f.BitLen() - 1
}

func maxTotalReward2(rewardValues []int) int {
	sort.Ints(rewardValues)
	rewardValues = slices.Compact(rewardValues)
	n := len(rewardValues)
	// f[i][r] 是否能在i位置能获得最大奖励r
	maxV := rewardValues[n-1]
	limit := 2 * maxV
	f := make([][]bool, n+1)
	for i := range f {
		f[i] = make([]bool, limit)
	}
	f[0][0] = true
	for i, v := range rewardValues {
		for j := range limit {
			if j >= v && j < 2*v {
				f[i+1][j] = f[i][j] || f[i][j-v]
			} else {
				f[i+1][j] = f[i][j]
			}
		}
	}
	ans := 0
	for j := 0; j < limit; j++ {
		if f[n][j] {
			ans = j
		}
	}
	return ans
}
