package _3_uniquePathsWithObstacles

// 63. 不同路径 II
// 给定一个 m x n 的整数数组 grid。一个机器人初始位于 左上角（即 grid[0][0]）。
// 机器人尝试移动到 右下角（即 grid[m - 1][n - 1]）。机器人每次只能向下或者向右移动一步。
//
// 网格中的障碍物和空位置分别用 1 和 0 来表示。机器人的移动路径中不能包含 任何 有障碍物的方格。
// 返回机器人能够到达右下角的不同路径数量。
//
// 测试用例保证答案小于等于 2 * 109。
//
// 示例 1：
// 输入：obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]]
//
// 输出：2
// 解释：3x3 网格的正中间有一个障碍物。
// 从左上角到右下角一共有 2 条不同的路径：
// 1. 向右 -> 向右 -> 向下 -> 向下
// 2. 向下 -> 向下 -> 向右 -> 向右
//
// 示例 2：
// 输入：obstacleGrid = [[0,1],[0,0]]
//
// 输出：1
// 提示：
//
// m == obstacleGrid.length
// n == obstacleGrid[i].length
// 1 <= m, n <= 100
// obstacleGrid[i][j] 为 0 或 1
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, m)
	for idx := range dp {
		dp[idx] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] == 0 && i > 0 {
			dp[i][0] = dp[i-1][0]
		} else {
			if obstacleGrid[i][0] == 1 {
				dp[i][0] = 0
			} else {
				dp[i][0] = 1
			}
		}
	}
	for j := 0; j < n; j++ {
		if obstacleGrid[0][j] == 0 && j > 0 {
			dp[0][j] = dp[0][j-1]
		} else {
			if obstacleGrid[0][j] == 1 {
				dp[0][j] = 0
			} else {
				dp[0][j] = 1
			}
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

// 使用滚动数组优化
// 逐行进行迭代, 每次迭代下一行的计算只依赖上一行
// 因此只需要定义一行的dp 即可
// 每迭代完一行了，下一行的计算通过累加自己来得到
func uniquePathsWithObstacles1(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([]int, n)
	if obstacleGrid[0][0] == 0 {
		dp[0] = 1
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[j] = 0
			} else if j >= 1 && obstacleGrid[i][j] == 0 {
				dp[j] += dp[j-1]
			}
		}
	}
	return dp[n-1]
}
