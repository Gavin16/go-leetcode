package _742_paintWalls

import (
	"math"
)

// 2742. 给墙壁刷油漆
// 给你两个长度为 n 下标从 0 开始的整数数组 cost 和 time ，
// 分别表示给 n 堵不同的墙刷油漆需要的开销和时间。你有两名油漆匠：
//
// 一位需要 付费 的油漆匠，刷第 i 堵墙需要花费 time[i] 单位的时间，开销为 cost[i] 单位的钱。
// 一位 免费 的油漆匠，刷 任意 一堵墙的时间为 1 单位，开销为 0 。但是必须在付费油漆匠 工作 时，免费油漆匠才会工作。
// 请你返回刷完 n 堵墙最少开销为多少。
//
// 示例 1：
//
// 输入：cost = [1,2,3,2], time = [1,2,3,2]
// 输出：3
// 解释：下标为 0 和 1 的墙由付费油漆匠来刷，需要 3 单位时间。同时，免费油漆匠刷下标为 2 和 3 的墙，需要 2 单位时间，开销为 0 。总开销为 1 + 2 = 3 。
// 示例 2：
//
// 输入：cost = [2,3,4,2], time = [1,1,1,1]
// 输出：4
// 解释：下标为 0 和 3 的墙由付费油漆匠来刷，需要 2 单位时间。同时，免费油漆匠刷下标为 1 和 2 的墙，需要 2 单位时间，开销为 0 。总开销为 2 + 2 = 4 。
//
// 提示：
//
// 1 <= cost.length <= 500
// cost.length == time.length
// 1 <= cost[i] <= 106
// 1 <= time[i] <= 500

func paintWalls(cost []int, time []int) int {
	n := len(cost)
	f := make([][]int, n+1)
	maxTm := 0
	for _, t := range time {
		maxTm = max(maxTm, t)
	}
	maxAdd := 2*max(n, maxTm) + 1
	for i := range f {
		f[i] = make([]int, maxAdd)
		for j := range f[i] {
			f[i][j] = math.MaxInt / 2
		}
	}
	f[0][0] = 0
	for i, c := range cost {
		tm := time[i]
		for j := maxAdd - 1; j >= 0; j-- {
			if j >= tm+1 {
				f[i+1][j] = min(f[i][j], f[i][j-1-tm]+c)
			} else {
				f[i+1][j] = f[i][j]
			}
		}
	}
	ans := math.MaxInt
	for j := maxAdd - 1; j >= n; j-- {
		ans = min(ans, f[n][j])
	}
	return ans
}

// 空间优化 -> 一维数组
// j 可以大于 n/2, 大于部分的状态保存到n/2 中
// 反推过程就变成了小于0的归的过来值保存到 0 中
// 这样做的好处就是: 超过n 的范围也能用 n 来做约束
func paintWalls1(cost []int, time []int) int {
	n := len(cost)
	f := make([]int, n+1)
	for i := range f {
		f[i] = math.MaxInt / 2
	}
	f[0] = 0
	for i, c := range cost {
		t := time[i] + 1
		for j := n; j > 0; j-- {
			f[j] = min(f[j], f[max(j-t, 0)]+c)
		}
	}
	return f[n]
}
