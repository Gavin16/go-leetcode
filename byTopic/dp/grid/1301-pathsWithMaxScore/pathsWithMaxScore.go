package _301_pathsWithMaxScore

import "slices"

// 1301. 最大得分的路径数目
// 给你一个正方形字符数组 board ，你从数组最右下方的字符 'S' 出发。
// 你的目标是到达数组最左上角的字符 'E' ，数组剩余的部分为数字字符 1, 2, ..., 9 或者障碍 'X'。
// 在每一步移动中，你可以向上、向左或者左上方移动，可以移动的前提是到达的格子没有障碍。
//
// 一条路径的 「得分」 定义为：路径上所有数字的和。
// 请你返回一个列表，包含两个整数：第一个整数是 「得分」 的最大值，第二个整数是得到最大得分的方案数，请把结果对 10^9 + 7 取余。
//
// 如果没有任何路径可以到达终点，请返回 [0, 0] 。
//
// 示例 1：
// 输入：board = ["E23","2X2","12S"]
// 输出：[7,1]
//
// 示例 2：
// 输入：board = ["E12","1X1","21S"]
// 输出：[4,2]
//
// 示例 3：
// 输入：board = ["E11","XXX","11S"]
// 输出：[0,0]
//
// 提示：
// 2 <= board.length == board[i].length <= 100

func pathsWithMaxScore(board []string) []int {
	mod := int(1e9 + 7)
	n := len(board)
	dp := make([][][2]int, n)
	for i := range dp {
		dp[i] = make([][2]int, n)
	}
	calScore := func(r, c int) int {
		if board[r][c] == 'S' {
			return 0
		}
		return int(board[r][c] - '0')
	}

	dp[0][0] = [2]int{0, 1}
	for i := 1; i < n; i++ {
		if board[0][i] == 'X' {
			dp[0][i] = [2]int{0, 0}
		} else if dp[0][i-1][0] > 0 || board[0][i-1] == 'E' {
			dp[0][i][0], dp[0][i][1] = dp[0][i-1][0]+calScore(0, i), dp[0][i-1][1]
		} else {
			dp[0][i] = dp[0][i-1]
		}
		if board[i][0] == 'X' {
			dp[i][0] = [2]int{0, 0}
		} else if dp[i-1][0][0] > 0 || board[i-1][0] == 'E' {
			dp[i][0][0], dp[i][0][1] = dp[i-1][0][0]+calScore(i, 0), dp[i-1][0][1]
		} else {
			dp[i][0] = dp[i-1][0]
		}
	}

	dirs := [][]int{{1, 0}, {0, 1}, {1, 1}}
	for i := 1; i < n; i++ {
		for j := 1; j < n; j++ {
			if board[i][j] == 'X' {
				dp[i][j] = [2]int{0, 0}
				continue
			}
			scoreDict := map[int]int{}
			scores := make([]int, 0)
			for _, dir := range dirs {
				li, lj := i-dir[0], j-dir[1]
				scoreDict[dp[li][lj][0]] = scoreDict[dp[li][lj][0]] + dp[li][lj][1]%mod
				scores = append(scores, dp[li][lj][0])
			}
			maxScore := slices.Max(scores)
			if maxScore == 0 && scoreDict[maxScore] == 0 {
				dp[i][j] = [2]int{0, 0}
			} else {
				dp[i][j] = [2]int{maxScore + calScore(i, j), scoreDict[maxScore]}
			}
		}
	}
	return []int{dp[n-1][n-1][0], dp[n-1][n-1][1]}
}
