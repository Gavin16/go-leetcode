package _186_maximumTotalDamage

import "slices"

// 3186. 施咒的最大总伤害
// 一个魔法师有许多不同的咒语。
// 给你一个数组 power ，其中每个元素表示一个咒语的伤害值，可能会有多个咒语有相同的伤害值。
// 已知魔法师使用伤害值为 power[i] 的咒语时，他们就 不能 使用伤害为 power[i] - 2 ，power[i] - 1 ，power[i] + 1 或者 power[i] + 2 的咒语。
// 每个咒语最多只能被使用 一次 。
// 请你返回这个魔法师可以达到的伤害值之和的 最大值 。
//
// 示例 1：
// 输入：power = [1,1,3,4]
// 输出：6
//
// 解释：
// 可以使用咒语 0，1，3，伤害值分别为 1，1，4，总伤害值为 6 。
//
// 示例 2：
// 输入：power = [7,1,6,6]
// 输出：13
// 解释：
// 可以使用咒语 1，2，3，伤害值分别为 1，6，6，总伤害值为 13 。
//
// 提示：
//
// 1 <= power.length <= 105
// 1 <= power[i] <= 109

func maximumTotalDamage(power []int) int64 {
	sums := make(map[int]int64)
	for _, p := range power {
		sums[p] = sums[p] + int64(p)
	}
	s := make([]int, 0)
	for k := range sums {
		s = append(s, k)
	}
	slices.Sort(s)
	j, n := 0, len(s)
	f := make([]int64, n+1)
	for i, x := range s {
		// 找出最后一个满足 s[j] < x-2的 元素下标
		for s[j] < x-2 {
			j++
		}
		f[i+1] = max(f[i], f[j]+sums[x])
	}
	return f[n]
}
