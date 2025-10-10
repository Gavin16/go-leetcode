package _2_minDistance

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	f := make([][]int, m+1)
	// f[0][j] 初始化为j, f[i][0] 初始化为 i
	for i, _ := range f {
		f[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		f[i][0] = i
	}
	for j := 1; j <= n; j++ {
		f[0][j] = j
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if word1[i] != word2[j] {
				f[i+1][j+1] = min(f[i][j], f[i+1][j], f[i][j+1]) + 1
			} else {
				f[i+1][j+1] = f[i][j]
			}
		}
	}
	return f[m][n]
}
