package _352_equalPairs

import "strings"

// 2352. 相等行列对
func equalPairs(grid [][]int) int {
	N := len(grid)
	map1, map2 := make(map[string]int), make(map[string]int)

	for i := 0; i < N; i++ {
		numStr := make([]string, 0, N)
		for j := 0; j < N; j++ {
			numStr = append(numStr, string(rune(grid[i][j])))
		}
		key := strings.Join(numStr, "-")
		map1[key] = map1[key] + 1
	}

	for j := 0; j < N; j++ {
		numStr := make([]string, 0, N)
		for i := 0; i < N; i++ {
			numStr = append(numStr, string(rune(grid[i][j])))
		}
		key := strings.Join(numStr, "-")
		map2[key] = map2[key] + 1
	}

	cnt := 0
	for k, v1 := range map1 {
		v2, ok := map2[k]
		if ok {
			cnt += v1 * v2
		}
	}
	return cnt
}
