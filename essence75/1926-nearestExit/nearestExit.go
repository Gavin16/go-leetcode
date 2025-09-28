package _926_nearestExit

// 1926. 迷宫中离入口最近的出口
// 给你一个 m x n 的迷宫矩阵 maze （下标从 0 开始），矩阵中有空格子（用 '.' 表示）和墙（用 '+' 表示）。
// 同时给你迷宫的入口 entrance ，用 entrance = [entrancerow, entrancecol] 表示你一开始所在格子的行和列。
//
// 每一步操作，你可以往 上，下，左 或者 右 移动一个格子。你不能进入墙所在的格子，你也不能离开迷宫。你的目标是找到离 entrance 最近 的出口。
// 出口 的含义是 maze 边界 上的 空格子。entrance 格子 不算 出口。
//
// 请你返回从 entrance 到最近出口的最短路径的 步数 ，如果不存在这样的路径，请你返回 -1 。
//
// 示例 1：
//
// 输入：maze = [["+","+",".","+"],[".",".",".","+"],["+","+","+","."]], entrance = [1,2]
// 输出：1
// 解释：总共有 3 个出口，分别位于 (1,0)，(0,2) 和 (2,3) 。
// 一开始，你在入口格子 (1,2) 处。
// - 你可以往左移动 2 步到达 (1,0) 。
// - 你可以往上移动 1 步到达 (0,2) 。
// 从入口处没法到达 (2,3) 。
// 所以，最近的出口是 (0,2) ，距离为 1 步。
// 示例 2：
//
// 输入：maze = [["+","+","+"],[".",".","."],["+","+","+"]], entrance = [1,0]
// 输出：2
// 解释：迷宫中只有 1 个出口，在 (1,2) 处。
// (1,0) 不算出口，因为它是入口格子。
// 初始时，你在入口与格子 (1,0) 处。
// - 你可以往右移动 2 步到达 (1,2) 处。
// 所以，最近的出口为 (1,2) ，距离为 2 步。
// 示例 3：
//
// 输入：maze = [[".","+"]], entrance = [0,0]
// 输出：-1
// 解释：这个迷宫中没有出口。
//
// 提示：
//
// maze.length == m
// maze[i].length == n
// 1 <= m, n <= 100
// maze[i][j] 要么是 '.' ，要么是 '+' 。
// entrance.length == 2
// 0 <= entrancerow < m
// 0 <= entrancecol < n
// entrance 一定是空格子。

func nearestExit(maze [][]byte, entrance []int) int {
	m, n := len(maze), len(maze[0])
	visited := make([][]byte, m)
	for id := range visited {
		visited[id] = make([]byte, n)
	}
	visited[entrance[0]][entrance[1]] = 1
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var nextGrids func(i, j int) [][]int
	nextGrids = func(i, j int) [][]int {
		nextSteps := make([][]int, 0)
		for _, dir := range dirs {
			ni, nj := i+dir[0], j+dir[1]
			if ni >= 0 && ni < m && nj >= 0 && nj < n &&
				maze[ni][nj] == '.' && visited[ni][nj] == 0 {
				nextSteps = append(nextSteps, []int{ni, nj})
				visited[ni][nj] = 1
			}
		}
		return nextSteps
	}
	stepCnt, grids := 0, nextGrids(entrance[0], entrance[1])
	for len(grids) > 0 {
		stepCnt = stepCnt + 1
		N := len(grids)
		for _, grid := range grids {
			row, col := grid[0], grid[1]
			if row == 0 || row == m-1 || col == 0 || col == n-1 {
				return stepCnt
			} else {
				next := nextGrids(row, col)
				grids = append(grids, next...)
			}
		}
		grids = grids[N:]
	}
	return -1
}

func nearestExit1(maze [][]byte, entrance []int) int {
	queue := make([][]int, 0)
	startR, startC := entrance[0], entrance[1]
	queue = append(queue, []int{startR, startC, 0})
	maze[startR][startC] = '+'

	m, n := len(maze), len(maze[0])
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for len(queue) > 0 {
		N := len(queue)
		for i := 0; i < N; i++ {
			r, c, d := queue[i][0], queue[i][1], queue[i][2]
			for _, dir := range dirs {
				nr, nc := r+dir[0], c+dir[1]
				if nr == -1 || nr == m || nc == -1 ||
					nc == n || maze[nr][nc] == '+' {
					continue
				}
				if nr == 0 || nr == m-1 || nc == 0 || nc == n-1 {
					return d + 1
				}
				queue = append(queue, []int{nr, nc, d + 1})
				maze[nr][nc] = '+'
			}
		}
		queue = queue[N:]
	}
	return -1
}
