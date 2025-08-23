package _6_merge

import "sort"

// !!!!!!!!!!!!!!!!!! 56. 合并区间  !!!!!!!!!!!!!!!!!!
// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
// 请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
//
// 示例 1：
//
// 输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
// 输出：[[1,6],[8,10],[15,18]]
// 解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
// 示例 2：
//
// 输入：intervals = [[1,4],[4,5]]
// 输出：[[1,5]]
// 解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。
//
// 提示：
//
// 1 <= intervals.length <= 104
// intervals[i].length == 2
// 0 <= starti <= endi <= 104
// TODO 考虑边界值

// 考虑边界条件
// 1. 排序后 后面的右边界小于前面的有边界 如 [1,4],[2,3]
// 2. 当存在大量被包裹的区域时，的判断条件 如 [1, 20],[2,3],[4,5],[6,9]
// 击败: 13.88%
func merge(intervals [][]int) [][]int {
	if len(intervals) == 1 {
		return [][]int{intervals[0]}
	}

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	minS, maxT, ans := intervals[0][0], intervals[0][1], make([][]int, 0)
	for i := 1; i < len(intervals); i += 1 {
		if intervals[i][0] > maxT {
			merged := []int{minS, maxT}
			minS, ans = intervals[i][0], append(ans, merged)
		}
		maxT = max(maxT, intervals[i][1])
	}
	return append(ans, []int{minS, maxT})
}
