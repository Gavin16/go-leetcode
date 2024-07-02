package _22_maxProfit2

// 122. 买卖股票的最佳时机 II
// 给你一个整数数组 prices ，其中 prices[i] 表示某支股票第 i 天的价格。
// 在每一天，你可以决定是否购买和/或出售股票。你在任何时候 最多 只能持有 一股 股票。你也可以先购买，然后在 同一天 出售。
//
// 返回 你能获得的 最大 利润 。
// 示例 1：
//
// 输入：prices = [7,1,5,3,6,4]
// 输出：7
// 解释：在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5 - 1 = 4。
// 随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6 - 3 = 3。
// 最大总利润为 4 + 3 = 7 。
// 示例 2：
//
// 输入：prices = [1,2,3,4,5]
// 输出：4
// 解释：在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5 - 1 = 4。
// 最大总利润为 4 。
// 示例 3：
//
// 输入：prices = [7,6,4,3,1]
// 输出：0
// 解释：在这种情况下, 交易无法获得正利润，所以不参与交易可以获得最大利润，最大利润为 0。
//
// 提示：
//
// 1 <= prices.length <= 3 * 104
// 0 <= prices[i] <= 104

// 子问题拆分:
// 采用动态规划算法
//  1. 定义状态
//  2. 确定状态转移方程
//  3. 初始化状态
//  4. 根据初始状态使用转移方程计算目标状态
//  5. 根据计算的目标状态得到结果返回

// 贪心算法
// 问题不限制交易次数，要计算最大的利润和，实际上是求 n 个区间的收益和
// 而对于单个区间来看 譬如（p1, pi] 区间的收益为 pi - p1 = (pi - pi-1) + (pi-1 - pi-2) + ... + (p2 - p1)
// 因此只要 pi - p1 收益为正(为负没必要进行交易), 就可以用相邻的区间元素差值累加来求得
// 相应的，如果求全局范围内相邻元素的差值, 通过筛选出大于0的差值，就能找出所有有收益的区间，这些区间的收益相加得到的就是问题所求
func maxProfit(prices []int) int {
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}

// 动态规划算法
// 题目要求每次最多只能持有"一股 股票" 定义状态转移方程对于任何一天都可能处于两种状态:持有股票 和 持有现金
// 定义 profit[i][0] 为第i天持有现金的状态，profit[i][1] 为第i天持有股票的状态
// 相应的状态转移方程为
// profit[i][0] = max(profit[i-1][1] + price[i], profit[i-1][0])
// profit[i][1] = max(profit[i-1][0] - price[i], profit[i-1][1])
// 对应的初始状态
// profit[0][0] = 0, profit[0][1] = -price[0]
func maxProfit1(prices []int) int {
	profit := make([][2]int, len(prices))
	profit[0][0] = 0
	profit[0][1] = -prices[0]

	for i := 1; i < len(prices); i++ {
		profit[i][0] = max(profit[i-1][1]+prices[i], profit[i-1][0])
		profit[i][1] = max(profit[i-1][0]-prices[i], profit[i-1][1])
	}
	return profit[len(prices)-1][0]
}