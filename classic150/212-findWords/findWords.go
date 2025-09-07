package _12_findWords

// 212. 单词搜索 II
// 给定一个 m x n 二维字符网格 board 和一个单词（字符串）列表 words， 返回所有二维网格上的单词。
//
// 单词必须按照字母顺序，通过 相邻的单元格 内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。
// 同一个单元格内的字母在一个单词中不允许被重复使用。
//
// 示例 1：
// 输入：board = [["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]],
//      words = ["oath","pea","eat","rain"]
//
// 输出：["eat","oath"]
// 示例 2：
//
// 输入：board = [["a","b"],["c","d"]], words = ["abcb"]
// 输出：[]
//
// 提示：
//
// m == board.length
// n == board[i].length
// 1 <= m, n <= 12
// board[i][j] 是一个小写英文字母
// 1 <= words.length <= 3 * 104
// 1 <= words[i].length <= 10
// words[i] 由小写英文字母组成
// words 中的所有字符串互不相同

// TODO 以 words 构建Trie 树，然后在 board 中搜索所有 word
// 时间复杂度: O(m*n*3^l) 空间复杂度: O(words.length)
// 击败:60.50%
func findWords(board [][]byte, words []string) []string {
	t := &Trie{}
	for _, word := range words {
		t.Insert(word)
	}
	m, n := len(board), len(board[0])
	found := map[string]bool{}

	var dfs func(node *Trie, x, y int)
	// 单次递归的时间复杂度为 3^l, 其中l为单词的长度
	// 因此除了board 中开始 tree 的初始节点，每个节点往下递归时，最多面临3个节点的后继节点
	dfs = func(node *Trie, x, y int) {
		ch := board[x][y]
		node = node.children[ch-'a']
		if node == nil {
			return
		}
		// 找到之后，后面可能还有单词,因此这里不能退出
		if node.word != "" {
			found[node.word] = true
		}
		board[x][y] = '#'
		for _, d := range dirs {
			nx, ny := x+d.x, y+d.y
			// 碰到边界 或者 深度搜索经过的路径 则退出; 否则继续搜索
			if inArea(nx, ny, m, n) && board[nx][ny] != '#' {
				dfs(node, nx, ny)
			}
		}
		board[x][y] = ch
	}
	for i, row := range board {
		for j := range row {
			dfs(t, i, j)
		}
	}
	ans := make([]string, 0)
	for s := range found {
		ans = append(ans, s)
	}
	return ans
}

var dirs = []struct{ x, y int }{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

type Trie struct {
	word     string
	children [26]*Trie
}

func (t *Trie) Insert(word string) {
	node := t
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &Trie{}
		}
		node = node.children[ch]
	}
	node.word = word
}

func inArea(i, j, m, n int) bool {
	if i < 0 || i >= m || j < 0 || j >= n {
		return false
	}
	return true
}
