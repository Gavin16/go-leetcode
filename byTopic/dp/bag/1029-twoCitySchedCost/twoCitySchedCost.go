package _029_twoCitySchedCost

import (
	"sort"
)

// 1029. 两地调度
// 公司计划面试 2n 人。给你一个数组 costs ，其中 costs[i] = [aCosti, bCosti] 。第 i 人飞往 a 市的费用为 aCosti ，飞往 b 市的费用为 bCosti 。
//
// 返回将每个人都飞到a 、b中某座城市的最低费用，要求每个城市都有 n 人抵达。
//
// 示例 1：
// 输入：costs = [[10,20],[30,200],[400,50],[30,20]]
// 输出：110
//
// 解释：
// 第一个人去 a 市，费用为 10。
// 第二个人去 a 市，费用为 30。
// 第三个人去 b 市，费用为 50。
// 第四个人去 b 市，费用为 20。
//
// 最低总费用为 10 + 30 + 50 + 20 = 110，每个城市都有一半的人在面试。
// 示例 2：
//
// 输入：costs = [[259,770],[448,54],[926,667],[184,139],[840,118],[577,469]]
// 输出：1859
// 示例 3：
//
// 输入：costs = [[515,563],[451,713],[537,709],[343,819],[855,779],[457,60],[650,359],[631,42]]
// 输出：3086
//
// 提示：
//
// 2 * n == costs.length
// 2 <= costs.length <= 100
// costs.length 为偶数
// 1 <= aCosti, bCosti <= 1000

func twoCitySchedCost(costs [][]int) int {
	s, n := 0, len(costs)/2
	for _, c := range costs {
		s += c[0]
	}
	diff := make([]int, 2*n)
	for i := range costs {
		diff[i] = costs[i][0] - costs[i][1]
	}
	sort.Slice(diff, func(i, j int) bool {
		return diff[i] > diff[j]
	})

	for i := 0; i < n; i++ {
		s -= diff[i]
	}
	return s
}
