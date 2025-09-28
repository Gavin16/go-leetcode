package _94_orangesRotting

// 994. 腐烂的橘子
// 在给定的 m x n 网格 grid 中，每个单元格可以有以下三个值之一：
//
// 值 0 代表空单元格；
// 值 1 代表新鲜橘子；
// 值 2 代表腐烂的橘子。
// 每分钟，腐烂的橘子 周围 4 个方向上相邻 的新鲜橘子都会腐烂。
//
// 返回 直到单元格中没有新鲜橘子为止所必须经过的最小分钟数。如果不可能，返回 -1 。
//
// 示例 1：
//
// 输入：grid = [[2,1,1],[1,1,0],[0,1,1]]
// 输出：4
// 示例 2：
//
// 输入：grid = [[2,1,1],[0,1,1],[1,0,1]]
// 输出：-1
// 解释：左下角的橘子（第 2 行， 第 0 列）永远不会腐烂，因为腐烂只会发生在 4 个方向上。
// 示例 3：
//
// 输入：grid = [[0,2]]
// 输出：0
// 解释：因为 0 分钟时已经没有新鲜橘子了，所以答案就是 0 。
//
// 提示：
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 10
// grid[i][j] 仅为 0、1 或 2

func orangesRotting(grid [][]int) int {
	rotten := make([][]int, 0)
	m, n, maxTm := len(grid), len(grid[0]), 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				rotten = append(rotten, []int{i, j})
			}
		}
	}
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for len(rotten) > 0 {
		N := len(rotten)
		for i := 0; i < N; i++ {
			ele := rotten[i]
			r, c := ele[0], ele[1]
			grid[r][c] = -2
			for _, dir := range dirs {
				nr, nc := r+dir[0], c+dir[1]
				if nr == -1 || nr == m || nc == -1 || nc == n ||
					grid[nr][nc] != 1 {
					continue
				}
				grid[nr][nc] = -2
				rotten = append(rotten, []int{nr, nc})
			}
		}
		rotten = rotten[N:]
		if len(rotten) > 0 {
			maxTm += 1
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}
	return maxTm
}
