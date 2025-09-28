package _431_kidsWithCandies

// 1431. 拥有最多糖果的孩子
// 有 n 个有糖果的孩子。给你一个数组 candies，其中 candies[i] 代表第 i 个孩子拥有的糖果数目，和一个整数 extraCandies 表示你所有的额外糖果的数量。
//
// 返回一个长度为 n 的布尔数组 result，如果把所有的 extraCandies 给第 i 个孩子之后，他会拥有所有孩子中 最多 的糖果，那么 result[i] 为 true，否则为 false。
//
// 注意，允许有多个孩子同时拥有 最多 的糖果数目。
//
// 示例 1：
//
// 输入：candies = [2,3,5,1,3], extraCandies = 3
// 输出：[true,true,true,false,true]
// 解释：如果你把额外的糖果全部给：
// 孩子 1，将有 2 + 3 = 5 个糖果，是孩子中最多的。
// 孩子 2，将有 3 + 3 = 6 个糖果，是孩子中最多的。
// 孩子 3，将有 5 + 3 = 8 个糖果，是孩子中最多的。
// 孩子 4，将有 1 + 3 = 4 个糖果，不是孩子中最多的。
// 孩子 5，将有 3 + 3 = 6 个糖果，是孩子中最多的。
// 示例 2：
//
// 输入：candies = [4,2,1,1,2], extraCandies = 1
// 输出：[true,false,false,false,false]
// 解释：只有 1 个额外糖果，所以不管额外糖果给谁，只有孩子 1 可以成为拥有糖果最多的孩子。
// 示例 3：
//
// 输入：candies = [12,1,12], extraCandies = 10
// 输出：[true,false,true]
//
// 提示：
//
// n == candies.length
// 2 <= n <= 100
// 1 <= candies[i] <= 100
// 1 <= extraCandies <= 50
func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxNum, N := 0, len(candies)
	for _, num := range candies {
		maxNum = max(maxNum, num)
	}
	ans := make([]bool, N)
	for i := 0; i < N; i++ {
		ans[i] = (extraCandies + candies[i]) >= maxNum
	}
	return ans
}
