package _2_totalNQueens

// 52. N 皇后 II
// n 皇后问题 研究的是如何将 n 个皇后放置在 n × n 的棋盘上，并且使皇后彼此之间不能相互攻击。
//
// 给你一个整数 n ，返回 n 皇后问题 不同的解决方案的数量。
//
// 示例 1：
//
// 输入：n = 4
// 输出：2
// 解释：如上图所示，4 皇后问题存在两个不同的解法。
// 示例 2：
//
// 输入：n = 1
// 输出：1
//
// 提示：
//
// 1 <= n <= 9

// 击败: 100.00%
func totalNQueens(n int) int {
	visit := make([][]bool, n)
	for i := range visit {
		visit[i] = make([]bool, n)
	}
	var dfs func(x int)

	count := 0
	dfs = func(x int) {
		if x == n {
			count += 1
			return
		}
		for i := 0; i < n; i++ {
			if positionOK(visit, x, i) {
				visit[x][i] = true
				dfs(x + 1)
				visit[x][i] = false
			}
		}
	}
	dfs(0)
	return count
}

// i,j 位置在行、列以及对角线方向上不能与其它元素相冲突
func positionOK(visit [][]bool, i, j int) bool {
	for k := 0; k < i; k++ {
		if visit[k][j] {
			return false
		}
	}
	for k := 0; k < j; k++ {
		if visit[k][j] {
			return false
		}
	}
	for k, l := i-1, j-1; k >= 0 && l >= 0; k, l = k-1, l-1 {
		if visit[k][l] {
			return false
		}
	}

	for k, l := i-1, j+1; k >= 0 && l < len(visit[0]); k, l = k-1, l+1 {
		if visit[k][l] {
			return false
		}
	}
	return true
}
