package _4_searchMatrix

// 74. 搜索二维矩阵
// 给你一个满足下述两条属性的 m x n 整数矩阵：
//
// 每行中的整数从左到右按非严格递增顺序排列。
// 每行的第一个整数大于前一行的最后一个整数。
// 给你一个整数 target ，如果 target 在矩阵中，返回 true ；否则，返回 false 。
//
// 示例 1：
//
// 输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
// 输出：true
// 示例 2：
//
// 输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
// 输出：false
//
// 提示：
//
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 100
// -104 <= matrix[i][j], target <= 104

// 该题放在二叉搜索的分类中，很容易让人想到利用矩阵的二叉结构来解题
// 实际上如果从左上角或者右下角的 位置来看，如果target < matrix[i][j] 那么只需要往都是小于 matrix[i][j]的一边去搜索即可
// 因为另一边的值大于 matrix[i][j], target 肯定不在其中
// 击败: 100.00%
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])

	i, j := m-1, 0
	for i >= 0 && j < n {
		if matrix[i][j] == target {
			return true
		} else if target < matrix[i][j] {
			i -= 1
		} else {
			j += 1
		}
	}
	return false
}
