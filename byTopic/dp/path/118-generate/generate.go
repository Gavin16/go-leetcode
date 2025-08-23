package _18_generate

// 118-杨辉三角形
// 给定一个非负整数 numRows，生成「杨辉三角」的前 numRows 行。
//
// 在「杨辉三角」中，每个数是它左上方和右上方的数的和。
//
// 示例 1:
//
// 输入: numRows = 5
// 输出: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
// 示例 2:
//
// 输入: numRows = 1
// 输出: [[1]]
//
// 提示:
//
// 1 <= numRows <= 30

// f[i][j] = f[i-1][j] + f[i-1][j-1]   j-1, j 在f[i-1] 的范围内
func generate(numRows int) [][]int {
	f := make([][]int, numRows)
	for idx := range f {
		f[idx] = make([]int, idx+1)
	}
	f[0][0] = 1

	for i := 1; i < numRows; i++ {
		for j := 0; j < i+1; j++ {
			if j == 0 || j == i {
				f[i][j] = 1
			} else {
				f[i][j] = f[i-1][j] + f[i-1][j-1]
			}
		}
	}
	return f
}
