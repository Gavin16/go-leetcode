package _16_canPartition

// 416. 分割等和子集
// 给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
//
// 示例 1：
//
// 输入：nums = [1,5,11,5]
// 输出：true
// 解释：数组可以分割成 [1, 5, 5] 和 [11] 。
// 示例 2：
//
// 输入：nums = [1,2,3,5]
// 输出：false
// 解释：数组不能分割成两个元素和相等的子集。
//
// 提示：
//
// 1 <= nums.length <= 200
// 1 <= nums[i] <= 100

func canPartition(nums []int) bool {
	sum, N := 0, len(nums)
	for _, v := range nums {
		sum += v
	}
	if N == 1 || sum%2 == 1 {
		return false
	}
	memo := make([]map[int]bool, N)
	for i := range memo {
		memo[i] = make(map[int]bool)
	}
	var dfs func(i int, left int) bool
	dfs = func(i int, left int) bool {
		if i >= N || left < 0 {
			return false
		}
		if left == 0 {
			return true
		}
		if can, ok := memo[i][left]; ok {
			return can
		}
		can := dfs(i+1, left-nums[i]) || dfs(i+1, left)
		memo[i][left] = can
		return can
	}
	return dfs(0, sum/2)
}

// f(i, left) 表示 0-i 范围内的数是否可以挑选出和为left 的组合
// 则 f(i, left) = f(i-1, left-nums[i]) || f(i-1, left)
// 初始值为f(-1, 0) = true
// 所求即为 f(n-1, sum/2)  其中 n = len(nums), sum = sum(nums)
func canPartition1(nums []int) bool {
	sum, n := 0, len(nums)
	for _, num := range nums {
		sum += num
	}
	if n == 1 || sum%2 == 1 {
		return false
	}
	f := make([][]bool, n+1)
	for i := range f {
		f[i] = make([]bool, (sum/2)+1)
	}
	f[0][0] = true
	for i := 0; i < n; i++ {
		for j := 0; j <= sum/2; j++ {
			f[i+1][j] = (j >= nums[i] && f[i][j-nums[i]]) || f[i][j]
		}
	}
	return f[n][sum/2]
}
