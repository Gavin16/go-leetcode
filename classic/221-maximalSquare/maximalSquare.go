package _21_maximalSquare

// 在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。
//
// 示例 1：
//
// 输入：matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
// 输出：4
// 示例 2：
//
// 输入：matrix = [["0","1"],["1","0"]]
// 输出：1
// 示例 3：
//
// 输入：matrix = [["0"]]
// 输出：0
//
// 提示：
//
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 300
// matrix[i][j] 为 '0' 或 '1'

// 定义f(i,j) 为以坐标(i-1,j-1) 为右下角的最大正方形的边长
// 则对于矩阵内任意的(i,j) i>0 且 j>0,
//
//			若matrix[i][j] = 0 则 f(i+1,j+1) = 0
//			若matrix[i][j] = 1 则 f(i+1, j+1) = 1 + min(f(i-1, j-1), f(i-1, j), f(i, j-1))
//
//		 对于f(1,1) 若matrix[0][0]=1 则 f(0,0) = 1, 否则f(0,0) = 0
//	  所求最大正方面的面积为 f(i,j) 对于所有(i,j) 中的最大值的平方

func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	f := make([][]int, m+1)
	for idx := range f {
		f[idx] = make([]int, n+1)
	}
	maxEdge := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '0' {
				f[i+1][j+1] = 0
			} else {
				f[i+1][j+1] = 1 + min(f[i][j], min(f[i][j+1], f[i+1][j]))
				maxEdge = max(maxEdge, f[i+1][j+1])
			}
		}
	}
	return maxEdge * maxEdge
}
