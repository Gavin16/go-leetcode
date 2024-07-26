package _27_ladderLength

// 127. 单词接龙
// 字典 wordList 中从单词 beginWord 到 endWord 的 转换序列 是一个按下述规格形成的序列 beginWord -> s1 -> s2 -> ... -> sk：
//
// 每一对相邻的单词只差一个字母。
// 对于 1 <= i <= k 时，每个 si 都在 wordList 中。注意， beginWord 不需要在 wordList 中。
// sk == endWord
// 给你两个单词 beginWord 和 endWord 和一个字典 wordList ，
// 返回 从 beginWord 到 endWord 的 最短转换序列 中的 单词数目 。如果不存在这样的转换序列，返回 0 。
//
// 示例 1：
//
// 输入：beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]
// 输出：5
// 解释：一个最短转换序列是 "hit" -> "hot" -> "dot" -> "dog" -> "cog", 返回它的长度 5。
// 示例 2：
//
// 输入：beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log"]
// 输出：0
// 解释：endWord "cog" 不在字典中，所以无法进行转换。
//
// 提示：
//
// 1 <= beginWord.length <= 10
// endWord.length == beginWord.length
// 1 <= wordList.length <= 5000
// wordList[i].length == beginWord.length
// beginWord、endWord 和 wordList[i] 由小写英文字母组成
// beginWord != endWord
// wordList 中的所有字符串 互不相同

// 击败: 20.44%
func ladderLength(beginWord string, endWord string, wordList []string) int {
	type Node struct {
		word   string
		adjoin []*Node
	}
	m := make(map[string]*Node)
	for _, word := range wordList {
		m[word] = &Node{word: word}
	}
	if _, ok := m[endWord]; !ok {
		return 0
	}
	minDist := len(wordList) + 1

	// 构造 wordList 对应的图 O(n^2)
	for _, v := range m {
		for _, w := range wordList {
			if wordDist(v.word, w) == 1 {
				nw := m[w]
				v.adjoin = append(v.adjoin, nw)
			}
		}
	}

	// 从与startWord 相邻单词列表开始搜索
	for _, w := range wordList {
		if wordDist(w, beginWord) != 1 {
			continue
		} else if w == endWord {
			return 2
		}

		visit := make(map[string]int)
		node, step := m[w], 1
		queue := []*Node{node}
	nearBFS:
		for len(queue) > 0 {
			size := len(queue)
			for i := 0; i < size; i++ {
				ni := queue[i]
				if visit[ni.word] == 1 {
					continue
				}
				visit[ni.word] = 1
				for _, adj := range ni.adjoin {
					if adj.word == endWord {
						minDist = min(minDist, step+1)
						break nearBFS
					} else if visit[adj.word] == 0 {
						queue = append(queue, adj)
					}
				}
			}
			queue, step = queue[size:], step+1
		}
	}
	if minDist > len(wordList) {
		return 0
	}
	// 转换序列长度 = 转换次数 + 1
	return minDist + 1
}

func wordDist(word1, word2 string) int {
	dist := 0
	for i := 0; i < len(word1); i++ {
		if word1[i] != word2[i] {
			dist += 1
		}
	}
	return dist
}
