package _23_maxProfit

import "slices"

// 123. 买卖股票的最佳时机 III
// 给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。
//
// 设计一个算法来计算你所能获取的最大利润。你最多可以完成 两笔 交易。
//
// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
//
// 示例 1:
//
// 输入：prices = [3,3,5,0,0,3,1,4]
// 输出：6
// 解释：在第 4 天（股票价格 = 0）的时候买入，在第 6 天（股票价格 = 3）的时候卖出，这笔交易所能获得利润 = 3-0 = 3 。
//
//	随后，在第 7 天（股票价格 = 1）的时候买入，在第 8 天 （股票价格 = 4）的时候卖出，这笔交易所能获得利润 = 4-1 = 3 。
//
// 示例 2：
//
// 输入：prices = [1,2,3,4,5]
// 输出：4
// 解释：在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
//
//	注意你不能在第 1 天和第 2 天接连购买股票，之后再将它们卖出。
//	因为这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。
//
// 示例 3：
//
// 输入：prices = [7,6,4,3,1]
// 输出：0
// 解释：在这个情况下, 没有交易完成, 所以最大利润为 0。
// 示例 4：
//
// 输入：prices = [1]
// 输出：0

// dp[i][0] 代表第i天 处于第一次买入状态时能获得的最大收益
// dp[i][1] 代表第i天 处于第一次卖出状态时能获得的最大收益
// dp[i][2] 代表第i天 处于第二次买入状态时能获得的最大收益
// dp[i][3] 代表第i天 处于第二次卖出状态时能获得的最大收益
// dp[i][0] = max(-price[i], dp[i-1][0])
// dp[i][1] = max(dp[i-1][1], dp[i-1][0] + prices[i])
// dp[i][2] = max(dp[i-1][2], dp[i-1][1] - prices[i])
// dp[i][3] = max(dp[i-1][3], dp[i-1][2] + prices[i])
func maxProfit(prices []int) int {
	N := len(prices)
	f := make([][]int, N)
	for i := range f {
		f[i] = make([]int, 4)
	}
	f[0][0], f[0][2] = -prices[0], -prices[0]
	for i := 1; i < N; i++ {
		price := prices[i]
		f[i][0] = max(-price, f[i-1][0])
		f[i][1] = max(f[i-1][1], f[i-1][0]+price)
		f[i][2] = max(f[i-1][2], f[i-1][1]-price)
		f[i][3] = max(f[i-1][3], f[i-1][2]+price)
	}
	return slices.Max(f[N-1])
}

func maxProfit1(prices []int) int {
	N := len(prices)
	f := make([]int, 4)
	f[0], f[2] = -prices[0], -prices[0]

	for i := 1; i < N; i++ {
		price := prices[i]
		f[3] = max(f[3], f[2]+price)
		f[2] = max(f[2], f[1]-price)
		f[1] = max(f[1], f[0]+price)
		f[0] = max(-price, f[0])
	}
	return slices.Max(f)
}
