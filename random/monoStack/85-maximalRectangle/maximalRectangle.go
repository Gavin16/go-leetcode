package _5_maximalRectangle

// 85-最大矩形
// 给定一个仅包含 0 和 1 、大小为 rows x cols 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。
//
// 示例 1：
//
// 输入：matrix = [
//
//					["1","0","1","0","0"],
//					["1","0","1","1","1"],
//					["1","1","1","1","1"],
//					["1","0","0","1","0"]
//	   			]
//
// 输出：6
// 解释：最大矩形如上图所示。
// 示例 2：
//
// 输入：matrix = [["0"]]
// 输出：0
// 示例 3：
//
// 输入：matrix = [["1"]]
// 输出：1
//
// 提示：
//
// rows == matrix.length
// cols == matrix[0].length
// 1 <= row, cols <= 200
// matrix[i][j] 为 '0' 或 '1'

// 使用单调栈解法
// 对于 matrix[i][j]  为'1' 的点, 找出该点第j列上侧(i-1 - 0) 范围内连续为'1'的点，记作up[i][j]
// 这样 up[i][j] 就会形成一个连续的纵向柱体, 通过使用单调栈 可 计算第i行所有柱体围起来的面积
// up[i][j] 范围内遍历所有行，每行都有一个面积最大的矩形，所有列面积最大的那个便是所求
func maximalRectangle(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	up := make([][]int, m)
	for idx := range up {
		up[idx] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '0' {
				up[i][j] = 0
				continue
			}
			up[i][j] = countUpOne(matrix, i, j)
		}
	}
	ans := 0
	for i := 0; i < m; i++ {
		leftIdx, rightIdx := make([]int, n), make([]int, n)
		leftIdx[0], rightIdx[n-1] = -1, n
		heights := up[i]
		for j := 1; j < n; j++ {
			temp := j - 1
			for temp >= 0 && heights[temp] >= heights[j] {
				temp = leftIdx[temp]
			}
			leftIdx[j] = temp
		}
		for k := n - 2; k >= 0; k-- {
			temp := k + 1
			for temp < n && heights[temp] >= heights[k] {
				temp = rightIdx[temp]
			}
			rightIdx[k] = temp
		}

		for s := 0; s < n; s++ {
			ans = max(ans, (rightIdx[s]-leftIdx[s]-1)*up[i][s])
		}
	}
	return ans
}

func countUpOne(matrix [][]byte, i, j int) int {
	count := 1
	for k := i - 1; k >= 0; k-- {
		if matrix[k][j] == '0' {
			break
		}
		count += 1
	}
	return count
}
