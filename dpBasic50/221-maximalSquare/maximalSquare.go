package _21_maximalSquare

func maximalSquare(matrix [][]byte) int {
	// f[i][j] 为以(i, j) 为右下角的最大正方形的边长
	m, n := len(matrix), len(matrix[0])
	f := make([][]int, m)
	for id := range f {
		f[id] = make([]int, n)
	}
	sideLen := 0
	for j := 0; j < n; j++ {
		f[0][j] = int(matrix[0][j] - '0')
		sideLen = max(sideLen, f[0][j])
	}
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '0' || j == 0 {
				f[i][j] = int(matrix[i][j] - '0')
				continue
			}
			f[i][j] = min(f[i-1][j-1], f[i-1][j], f[i][j-1]) + 1
			sideLen = max(sideLen, f[i][j])
		}
	}
	return sideLen * sideLen
}
