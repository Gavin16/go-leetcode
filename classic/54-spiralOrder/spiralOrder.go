package _4_spiralOrder

// 54. 螺旋矩阵
// 给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
//
// 示例 1：
//
// 输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
// 输出：[1,2,3,6,9,8,7,4,5]
// 示例 2：
//
// 输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
// 输出：[1,2,3,4,8,12,11,10,9,5,6,7]
//
// 提示：
//
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 10
// -100 <= matrix[i][j] <= 100

// 击败: 16.21%
func spiralOrder(matrix [][]int) []int {
	//顺时针螺旋 right -> down -> left -> up
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	i, j, did := 0, 0, 0
	rows, cols := len(matrix), len(matrix[0])

	ans := []int{matrix[0][0]}
	matrix[i][j] = -102
	for {
		dir := dirs[did]
		i, j = i+dir[0], j+dir[1]

		for i >= 0 && i < rows && j >= 0 && j < cols &&
			matrix[i][j] >= -100 {
			ans = append(ans, matrix[i][j])
			matrix[i][j] = -102
			i, j = i+dir[0], j+dir[1]
		}

		i, j, did = i-dir[0], j-dir[1], (did+1)%4
		dir = dirs[did]
		nextI, nextJ := i+dir[0], j+dir[1]
		if nextI < 0 || nextI >= rows || nextJ < 0 || nextJ >= cols ||
			matrix[nextI][nextJ] < -100 {
			break
		}
	}
	return ans
}
