package _76_findPaths

// 576. 出界的路径数
// 给你一个大小为 m x n 的网格和一个球。球的起始坐标为 [startRow, startColumn] 。
// 你可以将球移到在四个方向上相邻的单元格内（可以穿过网格边界到达网格之外）。你 最多 可以移动 maxMove 次球。
//
// 给你五个整数 m、n、maxMove、startRow 以及 startColumn ，找出并返回可以将球移出边界的路径数量。
// 因为答案可能非常大，返回对 109 + 7 取余 后的结果。
//
// 示例 1：
//
// 输入：m = 2, n = 2, maxMove = 2, startRow = 0, startColumn = 0
// 输出：6
// 示例 2：
//
// 输入：m = 1, n = 3, maxMove = 3, startRow = 0, startColumn = 1
// 输出：12
//
// 提示：
//
// 1 <= m, n <= 50
// 0 <= maxMove <= 50
// 0 <= startRow < m
// 0 <= startColumn < n

// 使用深度优先搜索路径
// f[i][j][k] 代表在(i,j)位置还剩k 步时的剩余的路径数量
// k == 0  则f[i][j][k] = 0
// k > 0  boundCnt=0;  若 k >= abs(m - i)  或者 k >= i+1  boundCnt += 1
//
//	若 k >= abs(n - j)  或者 k >= j+1   boundCnt += 1
//
// f[i][j][k] = f[i-1][j][k-1] + f[i+1][j][k-1] + f[i][j-1][k-1] + f[i][j+1][k-1] + boundCnt （i,j判断是否越界,根据情况调整）
// 返回 f[startRow][maxMove][maxMove]

// 对应动态规划解法只需要在一开始时对记忆化搜索的靠近边界的点做初始化
// 然后通过循环迭代 即可找出 f[startRow][startColumn][maxMove] 的值
// 循环以 (startRow, startColumn) 作为起始坐标点向四周扩散, 而不是顺序迭代
var mod = 1000000007

func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	f := make([][][]int, m)
	if maxMove == 0 {
		return 0
	}
	for rid := range f {
		f[rid] = make([][]int, n)
		for cid := range f[rid] {
			f[rid][cid] = make([]int, maxMove+1)
			for mm := range f[rid][cid] {
				f[rid][cid][mm] = -1
			}
		}
	}

	return dfs(f, m, n, maxMove, startRow, startColumn)
}

func countDirect(i, j, m, n int) int {
	count := 0
	if i == 0 {
		count += 1
	}
	if i == m-1 {
		count += 1
	}
	if j == 0 {
		count += 1
	}
	if j == n-1 {
		count += 1
	}
	return count
}

func dfs(f [][][]int, m, n, maxMv int, sr, sc int) int {
	if f[sr][sc][maxMv] != -1 {
		return f[sr][sc][maxMv]
	}
	if maxMv == 0 {
		f[sr][sc][maxMv] = 0
		return f[sr][sc][maxMv]
	}
	pathCount := countDirect(sr, sc, m, n)

	if sr > 0 {
		pathCount += dfs(f, m, n, maxMv-1, sr-1, sc) % mod
	}
	if sr < m-1 {
		pathCount += dfs(f, m, n, maxMv-1, sr+1, sc) % mod
	}
	if sc > 0 {
		pathCount += dfs(f, m, n, maxMv-1, sr, sc-1) % mod
	}
	if sc < n-1 {
		pathCount += dfs(f, m, n, maxMv-1, sr, sc+1) % mod
	}
	f[sr][sc][maxMv] = pathCount % mod
	return f[sr][sc][maxMv]
}

// 动态规划解法
func findPaths2(m int, n int, maxMove int, startRow int, startColumn int) int {
	f := make([][][]int, m)
	if maxMove == 0 {
		return 0
	}
	for rid := range f {
		f[rid] = make([][]int, n)
		for cid := range f[rid] {
			f[rid][cid] = make([]int, maxMove+1)
			for mm := range f[rid][cid] {
				if mm > 0 {
					count := 0
					if rid == 0 {
						count += 1
					}
					if rid == m-1 {
						count += 1
					}
					if cid == 0 {
						count += 1
					}
					if cid == n-1 {
						count += 1
					}
					f[rid][cid][mm] = count
				}
			}
		}
	}
	modVal := 1000000007

	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for step := 1; step <= maxMove; step++ {
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				for _, d := range dirs {
					ni, nj := i+d[0], j+d[1]
					if ni >= 0 && ni < m && nj >= 0 && nj < n {
						f[ni][nj][step] += f[i][j][step-1]
						f[ni][nj][step] %= modVal
					}
				}
			}
		}
	}
	return f[startRow][startColumn][maxMove]
}
