package _4_minPathSum

// 64. 最小路径和
// 给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
//
// 说明：每次只能向下或者向右移动一步。
//
// 示例 1：
//
// 输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
// 输出：7
// 解释：因为路径 1→3→1→1→1 的总和最小。
// 示例 2：
//
// 输入：grid = [[1,2,3],[4,5,6]]
// 输出：12
//
// 提示：
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 200
// 0 <= grid[i][j] <= 200

// 动态规划解法
// (1)定义dp[i][j] 代表 (i,j)位置的最短路径和
// (2)状态转移方程为:
//
//	dp[i][j] = min(dp[i-1][j] + grid[i][j], dp[i][j-1]+ grid[i][j])
//
// (3)初始条件 dp[0][0] = grid[0][0]
//
//	dp[0][j] = sum(grid[0][0], grid[0][1],..., grid[0][j])
//	dp[i][0] = sum(grid[0][0], grid[1][0], ...., grid[i][0])
//
// (4)要求的结果
//
//	全局最小路径和为： dp[m-1][n-1]
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for idx := range dp {
		dp[idx] = make([]int, n)
	}
	dp[0][0] = grid[0][0]
	for k := 1; k < n; k++ {
		dp[0][k] = dp[0][k-1] + grid[0][k]
	}
	for k := 1; k < m; k++ {
		dp[k][0] = dp[k-1][0] + grid[k][0]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[m-1][n-1]
}
