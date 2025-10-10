package _54_maxEnvelopes

import (
	"slices"
	"sort"
)

// 354. 俄罗斯套娃信封问题
// 给你一个二维整数数组 envelopes ，其中 envelopes[i] = [wi, hi] ，表示第 i 个信封的宽度和高度。
//
// 当另一个信封的宽度和高度都比这个信封大的时候，这个信封就可以放进另一个信封里，如同俄罗斯套娃一样。
//
// 请计算 最多能有多少个 信封能组成一组“俄罗斯套娃”信封（即可以把一个信封放到另一个信封里面）。
//
// 注意：不允许旋转信封。
//
// 示例 1：
//
// 输入：envelopes = [[5,4],[6,4],[6,7],[2,3]]
// 输出：3
// 解释：最多信封的个数为 3, 组合为: [2,3] => [5,4] => [6,7]。
// 示例 2：
//
// 输入：envelopes = [[1,1],[1,1],[1,1]]
// 输出：1
//
// 提示：
//
// 1 <= envelopes.length <= 105
// envelopes[i].length == 2
// 1 <= wi, hi <= 105

// O(n^2) 时间复杂度,对于数据规模 10^5 执行超时
func maxEnvelopes(envelopes [][]int) int {
	N := len(envelopes)
	sort.Slice(envelopes, func(i, j int) bool {
		return envelopes[i][0]*envelopes[i][1] <
			envelopes[j][0]*envelopes[j][1]
	})
	dp := make([]int, N)
	for i := 0; i < N; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if envelopes[i][0] > envelopes[j][0] &&
				envelopes[i][1] > envelopes[j][1] {
				dp[i] = max(dp[j]+1, dp[i])
			} else {
				dp[i] = max(dp[i], dp[j])
			}
		}
	}
	return slices.Max(dp)
}

// 将二维降低为1维数组
// 例如可以将 [[5,4],[6,4],[6,7],[2,3]] 处理为
// 2  5  6  6
// 3  4  7  4
// 这里当第一维数值相等时, 对第二维采用倒序排序,可以避免第二维错误的将 [6,4],[6,7] 都视为 可以套起来的信封
// 以上处理之后 就能忽略第一维的数值,直接处理第二维
// 而单独处理第二维的做法,就是求 最长递增子序列 LIS 问题
func maxEnvelopes1(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		a, b := envelopes[i], envelopes[j]
		return a[0] < b[0] || a[0] == b[0] && a[1] > b[1]
	})
	f := make([]int, 0)
	for _, e := range envelopes {
		h := e[1]
		if i := sort.SearchInts(f, h); i < len(f) {
			f[i] = h
		} else {
			f = append(f, h)
		}
	}
	return len(f)
}
