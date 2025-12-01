package _56_tallestBillboard

import "math"

// 956. 最高的广告牌
// 你正在安装一个广告牌，并希望它高度最大。这块广告牌将有两个钢制支架，两边各一个。每个钢支架的高度必须相等。
//
// 你有一堆可以焊接在一起的钢筋 rods。举个例子，如果钢筋的长度为 1、2 和 3，则可以将它们焊接在一起形成长度为 6 的支架。
//
// 返回 广告牌的最大可能安装高度 。如果没法安装广告牌，请返回 0 。
//
// 示例 1：
//
// 输入：[1,2,3,6]
// 输出：6
// 解释：我们有两个不相交的子集 {1,2,3} 和 {6}，它们具有相同的和 sum = 6。
// 示例 2：
//
// 输入：[1,2,3,4,5,6]
// 输出：10
// 解释：我们有两个不相交的子集 {2,3,5} 和 {4,6}，它们具有相同的和 sum = 10。
// 示例 3：
//
// 输入：[1,2]
// 输出：0
// 解释：没法安装广告牌，所以返回 0。
//
// 提示：
//
// 0 <= rods.length <= 20
// 1 <= rods[i] <= 1000
// sum(rods[i]) <= 5000

func tallestBillboard(rods []int) int {
	n, s := len(rods), 0
	for _, v := range rods {
		s += v
	}
	//limit := s / 2
	var dfs func(i, s1, s2 int) int
	dfs = func(i, s1, s2 int) int {
		if i < 0 {
			if s1 == s2 {
				return s1
			}
			return 0
		}
		if rods[i] > s1 || rods[i] > s2 {
			if rods[i] > s1 && rods[i] > s2 {
				return dfs(i-1, s1, s2)
			} else if rods[i] > s1 {
				return max(dfs(i-1, s1, s2+rods[i]), dfs(i-1, s1, s2))
			} else {
				return max(dfs(i-1, s1+rods[i], s2), dfs(i-1, s1, s2))
			}
		}
		return max(dfs(i-1, s1+rods[i], s2), dfs(i-1, s1, s2+rods[i]), dfs(i-1, s1, s2))
	}
	return dfs(n-1, 0, 0)
}

func tallestBillboard1(rods []int) int {
	n, s := len(rods), 0
	// dp[i][d] 代表以i结尾差值为d 组成一对杆中较短那个能取得的最大值
	for _, v := range rods {
		s += v
	}
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, s+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	dp[0][0] = 0
	for i, x := range rods {
		for d := 0; d <= s-x; d++ {
			short := dp[i][d]
			if short < 0 {
				continue
			}
			// 不选当前杆 x
			dp[i+1][d] = max(dp[i+1][d], dp[i][d])
			// 选择当前杆 x 放在较长的那一堆
			dp[i+1][d+x] = max(dp[i+1][d+x], dp[i][d])
			// 选择当前杆 x 放在较短的那一堆
			newDiff := int(math.Abs(float64(d - x)))
			dp[i+1][newDiff] = max(dp[i+1][newDiff], short+min(x, d))
		}
	}
	return dp[n][0]
}

func tallestBillboard2(rods []int) int {
	dp := make(map[int]int)
	dp[0] = 0
	for _, r := range rods {
		curr := map[int]int{}
		for k, v := range dp {
			curr[k] = v
		}
		for diff, h := range curr {
			// 选择将r 放到较短边
			newDiff1 := int(math.Abs(float64(diff - r)))
			shorter := min(h+r, h+diff)
			dp[newDiff1] = max(dp[newDiff1], shorter)

			// 选择将r 放到将长边
			newDiff2 := diff + r
			dp[newDiff2] = max(dp[newDiff2], h)
		}
	}
	return dp[0]
}
