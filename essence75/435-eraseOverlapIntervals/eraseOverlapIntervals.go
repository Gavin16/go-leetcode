package _35_eraseOverlapIntervals

import (
	"math"
	"sort"
)

// 435. 无重叠区间
// 给定一个区间的集合 intervals ，其中 intervals[i] = [starti, endi] 。
// 返回 需要移除区间的最小数量，使剩余区间互不重叠 。
//
// 注意 只在一点上接触的区间是 不重叠的。例如 [1, 2] 和 [2, 3] 是不重叠的。
//
// 示例 1:
//
// 输入: intervals = [[1,2],[2,3],[3,4],[1,3]]
// 输出: 1
// 解释: 移除 [1,3] 后，剩下的区间没有重叠。
// 示例 2:
//
// 输入: intervals = [ [1,2], [1,2], [1,2] ]
// 输出: 2
// 解释: 你需要移除两个 [1,2] 来使剩下的区间没有重叠。
// 示例 3:
//
// 输入: intervals = [ [1,2], [2,3] ]
// 输出: 0
// 解释: 你不需要移除任何区间，因为它们已经是无重叠的了。
//
// 提示:
//
// 1 <= intervals.length <= 105
// intervals[i].length == 2
// -5 * 104 <= starti < endi <= 5 * 104

// 找出最大不重叠的区间数量, 最小移除区间数 = 总数 - 最大不重叠数
func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	prevR := math.MinInt
	ans := 0
	for i := 0; i < len(intervals); i++ {
		if intervals[i][0] >= prevR {
			ans += 1
			prevR = intervals[i][1]
		}
	}
	return len(intervals) - ans
}
