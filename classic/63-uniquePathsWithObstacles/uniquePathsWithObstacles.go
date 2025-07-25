package _3_uniquePathsWithObstacles

// 63. 不同路径 II
// 给定一个 m x n 的整数数组 grid。一个机器人初始位于 左上角（即 grid[0][0]）。机器人尝试移动到 右下角（即 grid[m - 1][n - 1]）。机器人每次只能向下或者向右移动一步。
//
// 网格中的障碍物和空位置分别用 1 和 0 来表示。机器人的移动路径中不能包含 任何 有障碍物的方格。
//
// 返回机器人能够到达右下角的不同路径数量。
//
// 测试用例保证答案小于等于 2 * 109。
//
// 示例 1：
//
// 输入：obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]]
// 输出：2
// 解释：3x3 网格的正中间有一个障碍物。
// 从左上角到右下角一共有 2 条不同的路径：
// 1. 向右 -> 向右 -> 向下 -> 向下
// 2. 向下 -> 向下 -> 向右 -> 向右
// 示例 2：
//
// 输入：obstacleGrid = [[0,1],[0,0]]
// 输出：1
//
// 提示：
//
// m == obstacleGrid.length
// n == obstacleGrid[i].length
// 1 <= m, n <= 100
// obstacleGrid[i][j] 为 0 或 1

// 分析:  计算到达 dp[m][n] 的路径数量，满足dp[m][n] = dp[m-1][n] + dp[m][n-1]
// 动态规划解法
// (1) 定义f(i,j) 代表到达(i,j)位置的不同路径数量
// (2) 则对于 obstacleGrid[i][j]=0  f(i,j) = f(i-1, j) + f(i, j-1), 对于obstacleGrid[i][j]=1, 则f(i,j)=0
// (3) f(0,0) = 1, 对于f(0,j) j ∈ [1, n] 满足：
//
//		若 obstacleGrid[0][j] = 0  则 f(0, j) = f(0, j-1)
//		若 obstacleGrid[0][j] = 1  则 f(0, j) = 0
//	 类似的 f(i, 0) 也有这个规律
//
// (4) 返回结果为: f(i, j)
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	f := make([][]int, m)
	for id := range f {
		f[id] = make([]int, n)
	}

	f[0][0] = 1 - obstacleGrid[0][0]
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			pi, pj := i-1, j-1
			if obstacleGrid[i][j] == 0 {
				if pi >= 0 && pj >= 0 {
					f[i][j] = f[pi][j] + f[i][pj]
				} else {
					if pi >= 0 {
						f[i][j] = f[pi][j]
					} else {
						f[i][j] = f[i][pj]
					}
				}
			} else {
				f[i][j] = 0
			}

		}
	}
	return f[m-1][n-1]
}
