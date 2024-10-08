package _8_rotate

// 48. 旋转图像
// 给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。
//
// 你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。
//
// 示例 1：
//
// 输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
// 输出：[[7,4,1],[8,5,2],[9,6,3]]
// 示例 2：
//
// 输入：matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
// 输出：[[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]
//
// 提示：
//
// n == matrix.length == matrix[i].length
// 1 <= n <= 20
// -1000 <= matrix[i][j] <= 1000

// 遍历矩阵各个元素, 采用如下方式对矩阵元素进行轮转
// (i, j) -> (j, n-1-i)
// 遍历方式采用沿斜对角线元素作为起点 横向遍历 以 (n+1)/2 作为终点

// 击败: 32.29%
func rotate(matrix [][]int) {
	n := len(matrix)
	r := (n + 1) / 2

	for i := 0; i < r; i++ {
		for j := i; j < n-1-i; j++ {
			ci, cj, ni, nj := i, j, j, n-1-i
			cv := matrix[ci][cj]
			for k := 0; k < 4; k++ {
				nv := matrix[ni][nj]
				matrix[ni][nj] = cv
				ci, cj = ni, nj
				ni, nj = nj, n-1-ni
				cv = nv
			}
		}
	}
}
