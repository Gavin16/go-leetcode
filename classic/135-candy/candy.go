package _35_candy

import (
	"math"
)

// 135. 分发糖果
// n 个孩子站成一排。给你一个整数数组 ratings 表示每个孩子的评分。
//
// 你需要按照以下要求，给这些孩子分发糖果：
//
// 每个孩子至少分配到 1 个糖果。
// 相邻两个孩子评分更高的孩子会获得更多的糖果。
// 请你给每个孩子分发糖果，计算并返回需要准备的 最少糖果数目 。
//
// 示例 1：
//
// 输入：ratings = [1,0,2]
// 输出：5
// 解释：你可以分别给第一个、第二个、第三个孩子分发 2、1、2 颗糖果。
// 示例 2：
//
// 输入：ratings = [1,2,2]
// 输出：4
// 解释：你可以分别给第一个、第二个、第三个孩子分发 1、2、1 颗糖果。
//      第三个孩子只得到 1 颗糖果，这满足题面中的两个条件。
//
//
// 提示：
//
// n == ratings.length
// 1 <= n <= 2 * 104
// 0 <= ratings[i] <= 2 * 104

// 直观解法
// 计算ratings 的差分(第一个元素diff默认为0), 不为0的数值归一化
// 这时所有元素都变成了 0, 1, -1
// 对于取值 0 只有两种情况 ① 第一个元素 ② 该元素等于它左边的值
// 对于取值1 只有一种情况 它大于左边的值
// 对于取值-1 也只有一种情况 它小于左边的值
// 计算能分发的最少糖果数 可以根据如下情况来计算
// 对于diffs[i] 若满足
// (1) diffs[i] == 0, 左边的值要么是等于，要么是i==0不存在左边的元素
//
//	因此只需要看右边连续存在多少个-1, 是-1说明右边比自己少， 统计-1的个数 然后 candies[i] = count() + 1
//
// (2) diffs[i] == 1, 该位置的值大于左边，而右边的取值则不确定，可能等于，大于，也可能小于
//
//	只有大于右边的值的时候，才可能对当前位置的值有影响, 因此需要统计右边连续的-1数量
//	同时也统计左边连续1的数量， 最后取各边连续出现的数量的较大值  candies[i] = max(count(left) ,count(right)) + 1
//
// (3) diffs[i] == -1, 等于-1说明它比左边的值小，具体发多少糖果，需要以右边的为参考，统计右边连续-1出现次数即可
//
//	candies[i] = count(right) + 1
//
// 击败: 5.22%   --> 待优化
func candy(ratings []int) int {
	num := len(ratings)
	if num == 1 {
		return 1
	}
	diffs := make([]int, num)
	for i := 1; i < num; i++ {
		diffs[i] = ratings[i] - ratings[i-1]
		if diffs[i] != 0 {
			diffs[i] = diffs[i] / int(math.Abs(float64(diffs[i])))
		}
	}
	candies := make([]int, num)
	for i := 0; i < num; i++ {
		if diffs[i] == 1 {
			leftCnt, rightCnt := 0, 0
			for l, r := i, i+1; (l >= 0 && diffs[l] == 1) || (r < num && diffs[r] == -1); {
				if l >= 0 && diffs[l] == 1 {
					leftCnt++
					l--
				}
				if r < num && diffs[r] == -1 {
					rightCnt++
					r++
				}
			}
			candies[i] = max(leftCnt, rightCnt) + 1
		} else {
			rightCnt := 1
			for r := i + 1; r < num && diffs[r] == -1; r++ {
				if diffs[r] == -1 {
					rightCnt++
				}
			}
			candies[i] = rightCnt
		}
	}
	total := 0
	for _, v := range candies {
		total += v
	}
	return total
}

// 正向+反向 遍历累加
// 击败: 49.48%
func candy1(ratings []int) int {
	length := len(ratings)
	candies := make([]int, length)
	candies[0] = 1
	// 更新属于单调递增区间内的值
	for i := 1; i < length; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		} else {
			candies[i] = 1
		}
	}
	// length - 1 位置在正向更新时就能确定最终值
	// 更新单调递减区间内的值
	sum := candies[length-1]
	for j := length - 2; j >= 0; j-- {
		if ratings[j] > ratings[j+1] {
			candies[j] = max(candies[j], candies[j+1]+1)
		} else {
			candies[j] = max(candies[j], 1)
		}
		sum += candies[j]
	}
	return sum
}
